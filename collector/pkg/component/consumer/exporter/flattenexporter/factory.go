package flattenexporter

import (
	"fmt"
	"net/url"

	flattenTraces "github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/data/protogen/collector/trace/v1"
	flattenMetrics "github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/data/protogen/metrics/flatten"
	"github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/internal/component"
	"github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/internal/consumer"
	"github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/internal/exporterhelper"
)

func createTracesExporter(cfg *Cfg) (component.TracesExporter, error) {
	oCfg := cfg.Config
	return exporterhelper.NewTracesExporter(
		oCfg,
		cfg.Telemetry.GetZapLogger(),
		cfg.pushTraceData,
		exporterhelper.WithCapabilities(consumer.Capabilities{MutatesData: false}),
		exporterhelper.WithTimeout(exporterhelper.TimeoutSettings{Timeout: 0}),
		exporterhelper.WithRetry(oCfg.RetrySettings),
		exporterhelper.WithQueue(oCfg.QueueSettings))
}

func createMetricsExporter(cfg *Cfg) (component.MetricsExporter, error) {
	oCfg := cfg.Config
	return exporterhelper.NewMetricsExporter(
		oCfg,
		cfg.Telemetry.GetZapLogger(),
		cfg.pushMetricsData,
		exporterhelper.WithCapabilities(consumer.Capabilities{MutatesData: false}),
		exporterhelper.WithTimeout(exporterhelper.TimeoutSettings{Timeout: 0}),
		exporterhelper.WithRetry(oCfg.RetrySettings),
		exporterhelper.WithQueue(oCfg.QueueSettings))
}

func createTracesBatchProcessor(nextConsumer TracesBatch, cfg *Cfg, initBatchTraces flattenTraces.ExportTraceServiceRequest) (*Processor, error) {
	return NewBatchTracesProcessor(nextConsumer, cfg, initBatchTraces)
}

func createMetricsBatchProcessor(nextConsumer MetricsBatch, cfg *Cfg, initBatchMetrics flattenMetrics.FlattenMetrics) (*Processor, error) {
	return NewBatchMetricsProcessor(nextConsumer, cfg, initBatchMetrics)
}

func composeSignalURL(oCfg *Config, signalOverrideURL string, signalName string) (string, error) {
	switch {
	case signalOverrideURL != "":
		_, err := url.Parse(signalOverrideURL)
		if err != nil {
			return "", fmt.Errorf("%s_endpoint must be a valid URL", signalName)
		}
		return signalOverrideURL, nil
	case oCfg.Endpoint == "":
		return "", fmt.Errorf("either endpoint or %s_endpoint must be specified", signalName)
	default:
		endpoint := oCfg.Endpoint
		return endpoint + "/v1/" + signalName, nil
	}
}
