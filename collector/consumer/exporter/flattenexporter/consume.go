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
	service := e.Config.GetServiceInstance()
	switch gaugeGroup.Name {
	case constnames.SingleNetRequestGaugeGroup:
		singleTrace := transform.GenerateResourceSpans(gaugeGroup)
		traceServiceRequest := transform.CreateExportTraceServiceRequest(singleTrace)
		//to batchProcessor
		err := batchTraceProcessor.ConsumeTraces(context.Background(), traceServiceRequest)
		if err != nil {
			return err
		}
	case constnames.AggregatedNetRequestGaugeGroup:
		requestMetric := transform.GenerateRequestMetric(gaugeGroup)
		metricServiceRequest := transform.CreateFlattenMetrics(service, requestMetric)
		err := batchMetricProcessor.ConsumeMetrics(context.Background(), metricServiceRequest)
		if err != nil {
			return err
		}
		//TCP 链接指标
	case constnames.TcpStatsGaugeGroup:
		tcpInuseMetric := transform.GenerateXXMetric(gaugeGroup, constant.MetricTypeTcpInuse)
		metricServiceRequest := transform.CreateFlattenMetrics(service, tcpInuseMetric)
		err := batchMetricProcessor.ConsumeMetrics(context.Background(), metricServiceRequest)
		if err != nil {
			return err
		}
	case constnames.PageFaultGaugeGroupName:
		pageFaultMetric := transform.GenerateXXMetric(gaugeGroup, constant.MetricTypePageFault)
		metricServiceRequest := transform.CreateFlattenMetrics(service, pageFaultMetric)
		err := batchMetricProcessor.ConsumeMetrics(context.Background(), metricServiceRequest)
		if err != nil {
			return err
		}
	case constnames.ConnectGaugeGroupName:
		connectMetric := transform.GenerateXXMetric(gaugeGroup, constant.MetricTypeConnect)
		metricServiceRequest := transform.CreateFlattenMetrics(service, connectMetric)
		err := batchMetricProcessor.ConsumeMetrics(context.Background(), metricServiceRequest)
		if err != nil {
			return err
		}
	default:
		return nil
	}
	return nil
}
