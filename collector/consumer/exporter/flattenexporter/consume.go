package flattenexporter

import (
	"context"
	"github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/constant"
	"github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/transform"
	"github.com/Kindling-project/kindling/collector/model"
	"github.com/Kindling-project/kindling/collector/model/constnames"
)

func (e *Cfg) Consume(gaugeGroup *model.GaugeGroup) error {
	if gaugeGroup == nil {
		return nil
	}

	if e.pushExporters == nil {
		return nil
	}
	if e.batchProcessors == nil {
		return nil
	}

	batchTraceProcessor := (*e.batchProcessors)[constant.Traces]
	batchMetricProcessor := (*e.batchProcessors)[constant.Metrics]
	switch gaugeGroup.Name {
	case constnames.SingleNetRequestGaugeGroup:
		traceServiceRequest := transform.CreateExportTraceServiceRequest(transform.GenerateResourceSpans(gaugeGroup))
		//to batchProcessor
		err := batchTraceProcessor.ConsumeTraces(context.Background(), traceServiceRequest)
		if err != nil {
			return err
		}
	case constnames.AggregatedNetRequestGaugeGroup:
		traceServiceRequest := transform.CreateFlattenMetrics(e.Config.GetServiceInstance(), transform.GenerateMetrics(gaugeGroup))
		err := batchMetricProcessor.ConsumeMetrics(context.Background(), traceServiceRequest)
		if err != nil {
			return err
		}
	default:
		return nil
	}
	return nil
}
