package flattenexporter

import (
	"context"
	"github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/data/protogen/collector/metrics/flatten"
	exportTrace "github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/data/protogen/collector/trace/flatten"
	flattenTraces "github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/data/protogen/collector/trace/v1"
	flattenMetrics "github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/data/protogen/metrics/flatten"
	"github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/internal/component"
	"go.uber.org/zap"
	"strconv"
)

type MetricsBatch interface {
	ConsumeMetrics(ctx context.Context, md flattenMetrics.FlattenMetrics) error
}

type TracesBatch interface {
	ConsumeTraces(ctx context.Context, td flattenTraces.ExportTraceServiceRequest) error
}

type Consumer struct {
	TracesExporter  component.TracesExporter
	MetricsExporter component.MetricsExporter
	cfg             *Cfg
}

func (e *Consumer) ConsumeTraces(context context.Context, td flattenTraces.ExportTraceServiceRequest) error {
	tracesRequest := exportTrace.ExportFlattenTraceServiceRequest{
		Service:            e.cfg.Config.GetServiceInstance(),
		ResourceSpansBytes: make([][]byte, len(td.ResourceSpans)),
	}
	if ce := e.cfg.Telemetry.Logger.Check(zap.DebugLevel, "ExportTraceServiceRequest"); ce != nil {
		ce.Write(
			zap.String("size", strconv.Itoa(len(td.ResourceSpans))))
	}
	for i := 0; i < len(td.ResourceSpans); i++ {
		resourceSpanBytes, err := td.ResourceSpans[i].Marshal()
		if err != nil {
			continue
		} else {
			tracesRequest.ResourceSpansBytes[i] = resourceSpanBytes
		}
	}
	marshal, err := tracesRequest.Marshal()
	if err != nil {
		e.cfg.Telemetry.Logger.Error("tracesRequest marshal fail", zap.Error(err))
		return err
	}
	return e.TracesExporter.ConsumeTraces(context, marshal)
}

func (e *Consumer) ConsumeMetrics(context context.Context, md flattenMetrics.FlattenMetrics) error {
	if ce := e.cfg.Telemetry.Logger.Check(zap.DebugLevel, "FlattenMetrics"); ce != nil {
		ce.Write(
			zap.String("size", strconv.Itoa(len(md.RequestMetricByte.Metrics))))
	}
	requestMetricsBytes, requestErr := md.RequestMetricByte.Marshal()
	if requestErr != nil {
		return requestErr
	}
	metricsRequest := &flatten.FlattenExportMetricsServiceRequest{
		RequestMetrics: requestMetricsBytes,
		Service:        e.cfg.Config.GetServiceInstance(),
	}

	marshal, err := metricsRequest.Marshal()
	if err != nil {
		e.cfg.Telemetry.Logger.Error("metricsRequest marshal fail", zap.Error(err))
		return err
	}
	return e.MetricsExporter.ConsumeMetrics(context, marshal)
}
