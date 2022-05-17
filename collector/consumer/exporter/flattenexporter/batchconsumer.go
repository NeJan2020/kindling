package flattenexporter

import (
	"context"
	"fmt"
	"github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/data/protogen/collector/metrics/flatten"
	exportTrace "github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/data/protogen/collector/trace/flatten"
	flattenTraces "github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/data/protogen/collector/trace/v1"
	flattenMetrics "github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/data/protogen/metrics/flatten"
	"github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/internal/component"
	"github.com/Kindling-project/kindling/collector/model/constvalues"
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
	for i := 0; i < len(td.ResourceSpans); i++ {
		for j := 0; j < len(td.ResourceSpans[i].InstrumentationLibrarySpans); j++ {
			fmt.Println(td.ResourceSpans[i].InstrumentationLibrarySpans[j].Spans[0].Events[0].Attributes[30].Value.GetValue())
		}
	}
	tracesRequest := exportTrace.ExportFlattenTraceServiceRequest{
		Service:            e.cfg.Config.GetServiceInstance(),
		ResourceSpansBytes: make([][]byte, len(td.ResourceSpans)),
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
		return err
	}
	return e.TracesExporter.ConsumeTraces(context, marshal)
}

func (e *Consumer) ConsumeMetrics(context context.Context, md flattenMetrics.FlattenMetrics) error {
	for i := 0; i < len(md.RequestMetrics.Metrics); i++ {
		fmt.Println(md.RequestMetrics.Metrics[i].GetMetricMap()[constvalues.RequestTotalTime].GetSum().GetValue())
	}
	requestMetricsBytes, err := md.RequestMetrics.Marshal()
	if err != nil {
		return err
	}
	metricsRequest := &flatten.FlattenExportMetricsServiceRequest{
		RequestMetrics:   requestMetricsBytes,
		ConnectMetrics:   nil,
		InterfaceMetrics: nil,
		Service:          e.cfg.Config.GetServiceInstance(),
	}

	marshal, err := metricsRequest.Marshal()
	if err != nil {
		return err
	}
	return e.MetricsExporter.ConsumeMetrics(context, marshal)
}
