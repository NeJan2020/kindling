package elasticexporter

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/Kindling-project/kindling/collector/pkg/component"
	"github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter"
	"github.com/Kindling-project/kindling/collector/pkg/model"
	es7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"go.uber.org/zap"
)

const DefaultTraceIndex = "kindling-trace-test-v0.0.1"
const DefaultMetricIndex = "kindling-metric-test-v0.0.1"
const Type = "elasticexporter"

type ElasticExporter struct {
	client    *es7.Client
	cfg       *Config
	telemetry *component.TelemetryTools
}

func FormatJson(data string) string {
	var out bytes.Buffer
	json.Indent(&out, []byte(data), "", "\t")
	return out.String()
}

func New(config interface{}, telemetry *component.TelemetryTools) exporter.Exporter {
	var cfg *Config
	var ok bool
	if cfg, ok = config.(*Config); !ok {
		if cfg.ESConfig == nil {
			telemetry.Logger.Error("failed to find es-config in [ exporter.elasticexporter ], can not build es")
		}
		telemetry.Logger.Error("elatsicSearch Exporter Config Paresed failed!")
		return nil
	}
	client, err := es7.NewClient(*cfg.ESConfig)
	if err != nil {
		telemetry.Logger.Error("unexpected es-config!", zap.Error(err))
		return nil
	}
	return &ElasticExporter{
		cfg:       cfg,
		telemetry: telemetry,
		client:    client,
	}
}

func (e *ElasticExporter) AddTraceToBulk(trace *model.DataGroup) {
	// TODO implement me
}

func (e *ElasticExporter) exportTrace(trace *model.DataGroup) error {
	if e.cfg.EnableBulk {
		e.AddTraceToBulk(trace)
	}
	data, err := json.Marshal(trace)
	if err != nil {
		e.telemetry.Logger.Error("failed to marshal trace into json", zap.String("data", string(data)))
		return err
	}
	req := esapi.IndexRequest{
		Index: e.cfg.ESIndex.TraceIndex,
		Body:  bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), e.client)
	if err != nil {
		e.telemetry.Logger.Error("failed to send index req to elasticsearch", zap.Error(err), zap.String("data", string(data)))
		return err
	}
	defer res.Body.Close()
	return nil
}

func (e *ElasticExporter) exportMetric(metric *model.DataGroup) error {
	if e.cfg.EnableBulk {
		e.AddTraceToBulk(metric)
	}
	data, err := json.Marshal(metric)
	if err != nil {
		e.telemetry.Logger.Error("failed to marshal metric into json", zap.String("metric", metric.String()))
		return err
	}
	req := esapi.IndexRequest{
		Index: e.cfg.ESIndex.MetricIndex,
		Body:  bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), e.client)
	if err != nil {
		e.telemetry.Logger.Error("failed to send index req to elasticsearch", zap.Error(err), zap.String("data", string(data)))
		return err
	}
	defer res.Body.Close()
	return nil
}
