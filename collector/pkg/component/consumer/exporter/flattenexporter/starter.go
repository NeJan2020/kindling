package flattenexporter

import (
	"context"

	"github.com/Kindling-project/kindling/collector/pkg/component"
	"github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/constant"
	flattenTraces "github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/data/protogen/collector/trace/v1"
	v1 "github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/data/protogen/common/v1"
	flattenMetrics "github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/data/protogen/metrics/flatten"
	trace "github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/data/protogen/trace/v1"
	internalComponent "github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/internal/component"
	"github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/transform"
	"go.uber.org/zap"
)

func createAndStartRetryExporter(oce *Cfg, telemetry *component.TelemetryTools) *map[string]internalComponent.Exporter {
	exporters := make(map[string]internalComponent.Exporter, 0)
	tracesExporter, tracesExporterErr := createTracesExporter(oce)
	metricsExporter, metricsExporterErr := createMetricsExporter(oce)

	if tracesExporterErr != nil {
		telemetry.Logger.Panic("Cannot createTracesExporters", zap.String("componentType", Flattten))
		return nil
	}
	tracesExporterStartErr := tracesExporter.Start(context.Background(), nil)
	if tracesExporterStartErr != nil {
		telemetry.Logger.Panic("Cannot start tracesExporter", zap.String("componentType", Flattten))
		return nil
	}

	if metricsExporterErr != nil {
		telemetry.Logger.Panic("Cannot createMetricsExporter", zap.String("componentType", Flattten))
		return nil
	}
	metricsExporterStartErr := metricsExporter.Start(context.Background(), nil)
	if metricsExporterStartErr != nil {
		telemetry.Logger.Panic("Cannot start metricsExporter", zap.String("componentType", Flattten))
		return nil
	}
	exporters[constant.Traces] = tracesExporter
	exporters[constant.Metrics] = metricsExporter
	return &exporters
}

func InitTraceRequest() flattenTraces.ExportTraceServiceRequest {
	return transform.CreateExportTraceServiceRequest(make([]*trace.ResourceSpans, 0))
}

func InitMetricsRequest(service *v1.Service) flattenMetrics.FlattenMetrics {
	return transform.CreateFlattenMetrics(service, make([]*flattenMetrics.RequestMetric, 0))
}

func createAndStartBatchProcessor(oce *Cfg, telemetry *component.TelemetryTools) *map[string]*Processor {
	batchProcessors := make(map[string]*Processor, 0)

	pushTraceExporter := (*oce.pushExporters)[constant.Traces].(internalComponent.TracesExporter)
	pushMetricsExporter := (*oce.pushExporters)[constant.Metrics].(internalComponent.MetricsExporter)

	batchConsumer := &Consumer{TracesExporter: pushTraceExporter, MetricsExporter: pushMetricsExporter, cfg: oce}

	//create and start TracesBatchProcessor
	tracesBatchProcessor, tracesBatchErr := createTracesBatchProcessor(batchConsumer, oce, InitTraceRequest())
	if tracesBatchErr != nil {
		telemetry.Logger.Panic("Cannot createTracesBatchProcessor", zap.String("componentType", Flattten))
		return nil
	}
	tracesBatchErr = tracesBatchProcessor.Start(context.Background())
	if tracesBatchErr != nil {
		return nil
	}
	//create and start MetricsBatchProcessor
	metricsBatchProcessor, metricsBatchErr := createMetricsBatchProcessor(batchConsumer, oce, InitMetricsRequest(oce.Config.GetServiceInstance()))
	if metricsBatchErr != nil {
		telemetry.Logger.Panic("Cannot createMetricsBatchProcessor", zap.String("componentType", Flattten))
		return nil
	}
	metricsBatchErr = metricsBatchProcessor.Start(context.Background())
	if metricsBatchErr != nil {
		return nil
	}

	batchProcessors[constant.Traces] = tracesBatchProcessor
	batchProcessors[constant.Metrics] = metricsBatchProcessor
	return &batchProcessors
}
