package flattenexporter

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/Kindling-project/kindling/collector/pkg/component"
	"github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter"
	"github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/constant"
	internalComponent "github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/internal/component"
	"github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/internal/consumer/consumererror"
	"github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/internal/exporterhelper"
	"github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/internal/middleware"
	"github.com/Kindling-project/kindling/collector/pkg/processpid"

	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/proto"
)

type Cfg struct {
	// Input configuration.
	Config          *Config
	client          *http.Client
	Telemetry       *component.TelemetryTools
	pushExporters   *map[string]internalComponent.Exporter
	batchProcessors *map[string]*Processor
}

const (
	headerRetryAfter         = "Retry-After"
	maxHTTPResponseReadBytes = 64 * 1024
	Flattten                 = "flattenexporter"
)

var initProcessPidPortInfo sync.Once

func NewExporter(config interface{}, telemetry *component.TelemetryTools) exporter.Exporter {
	telemetry.Logger.Info("Create Flatten Exporter!")
	oce, err := newCfg(config, telemetry)
	initProcessPidPortInfo.Do(func() {
		config := oce.Config
		// Try to get HostIP from env
		hostIp := os.Getenv("MY_NODE_IP")
		if hostIp == "" {
			hostIp = "UNKNOW_HOST"
		}
		telemetry.Logger.Info("FlattenExporter is creating ProcessPid Fetch Task!")
		processpid.InitSendPidPortBytime(config.BatchTimeout, config.Endpoint, config.MasterIp, hostIp, telemetry)
	})
	if err != nil {
		telemetry.Logger.Panic("Cannot convert Component config", zap.String("componentType", Flattten))
		return nil
	}
	oce.pushExporters = createAndStartRetryExporter(oce, telemetry)
	oce.batchProcessors = createAndStartBatchProcessor(oce, telemetry)
	return oce
}

// Crete Cfg.
func newCfg(cfg interface{}, telemetry *component.TelemetryTools) (*Cfg, error) {
	oCfg := cfg.(*Config)
	if oCfg.Endpoint != "" {
		_, err := url.Parse(oCfg.Endpoint)
		if err != nil {
			return nil, errors.New("endpoint must be a valid URL")
		}
	}

	client, err := oCfg.HTTPClientSettings.ToClient()
	if err != nil {
		return nil, err
	}

	if oCfg.Compression != "" {
		if strings.ToLower(oCfg.Compression) == "gzip" {
			client.Transport = middleware.NewCompressRoundTripper(client.Transport)
		} else {
			return nil, fmt.Errorf("unsupported compression type %q", oCfg.Compression)
		}
	}

	return &Cfg{
		Config:    oCfg,
		client:    client,
		Telemetry: telemetry,
	}, nil
}

func (e *Cfg) pushTraceData(ctx context.Context, request []byte) error {
	traceUrl, _ := composeSignalURL(e.Config, e.Config.TracesEndpoint, constant.Traces)
	return e.httpExport(ctx, traceUrl, request)
}

func (e *Cfg) pushMetricsData(ctx context.Context, request []byte) error {
	metricUrl, _ := composeSignalURL(e.Config, e.Config.MetricsEndpoint, constant.Metrics)
	return e.httpExport(ctx, metricUrl, request)
}

func (e *Cfg) httpExport(ctx context.Context, url string, request []byte) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(request))
	if err != nil {
		return consumererror.Permanent(err)
	}
	req.Header.Set("Content-Type", "application/x-protobuf")
	resp, err := e.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make an HTTP request to endpoint %s: %w", url, err)
	}

	defer func() {
		// Discard any remaining response body when we are done reading.
		io.CopyN(ioutil.Discard, resp.Body, maxHTTPResponseReadBytes) // nolint:errcheck
		resp.Body.Close()
	}()

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		// Request is successful.
		return nil
	}

	respStatus := readResponse(resp)

	// Format the error message. Use the status if it is present in the response.
	var formattedErr error
	if respStatus != nil {
		formattedErr = fmt.Errorf(
			"error exporting items, request to %s responded with HTTP Status Code %d, Message=%s, Details=%v",
			url, resp.StatusCode, respStatus.Message, respStatus.Details)
	} else {
		formattedErr = fmt.Errorf(
			"error exporting items, request to %s responded with HTTP Status Code %d",
			url, resp.StatusCode)
	}

	// Check if the server is overwhelmed.
	// See spec https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/protocol/otlp.md#throttling-1
	if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode == http.StatusServiceUnavailable {
		// Fallback to 0 if the Retry-After header is not present. This will trigger the
		// default backoff policy by our caller (retry handler).
		retryAfter := 0
		if val := resp.Header.Get(headerRetryAfter); val != "" {
			if seconds, err2 := strconv.Atoi(val); err2 == nil {
				retryAfter = seconds
			}
		}
		// Indicate to our caller to pause for the specified number of seconds.
		return exporterhelper.NewThrottleRetry(formattedErr, time.Duration(retryAfter)*time.Second)
	}

	if resp.StatusCode == http.StatusBadRequest {
		// Report the failure as permanent if the server thinks the request is malformed.
		return consumererror.Permanent(formattedErr)
	}

	// All other errors are retryable, so don't wrap them in consumererror.Permanent().
	return formattedErr
}

// Read the response and decode the status.Status from the body.
// Returns nil if the response is empty or cannot be decoded.
func readResponse(resp *http.Response) *status.Status {
	var respStatus *status.Status
	if resp.StatusCode >= 400 && resp.StatusCode <= 599 {
		// Request failed. Read the body. OTLP spec says:
		// "Response body for all HTTP 4xx and HTTP 5xx responses MUST be a
		// Protobuf-encoded Status message that describes the problem."
		maxRead := resp.ContentLength
		if maxRead == -1 || maxRead > maxHTTPResponseReadBytes {
			maxRead = maxHTTPResponseReadBytes
		}
		respBytes := make([]byte, maxRead)
		n, err := io.ReadFull(resp.Body, respBytes)
		if err == nil && n > 0 {
			// Decode it as Status struct. See https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/protocol/otlp.md#failures
			respStatus = &status.Status{}
			err = proto.Unmarshal(respBytes, respStatus)
			if err != nil {
				respStatus = nil
			}
		}
	}

	return respStatus
}
