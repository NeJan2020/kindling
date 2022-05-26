package flattenexporter

import (
	"context"
	"go.uber.org/zap"

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
	e.Telemetry.Logger.Info("GaugeGroup...." + gaugeGroup.String())
	service := e.Config.GetServiceInstance()
	switch gaugeGroup.Name {
	case constnames.SingleNetRequestGaugeGroup:
		singleTrace := transform.GenerateResourceSpans(gaugeGroup)
		traceServiceRequest := transform.CreateExportTraceServiceRequest(singleTrace)
		err := batchTraceProcessor.ConsumeTraces(context.Background(), traceServiceRequest)
		if err != nil {
			e.Telemetry.Logger.Error("Failed to consume metrics single_net_request_gauge_group", zap.Error(err))
			return err
		}
	case constnames.AggregatedNetRequestGaugeGroup:
		requestMetric := transform.GenerateRequestMetric(gaugeGroup)
		metricServiceRequest := transform.CreateFlattenMetrics(service, requestMetric)
		err := batchMetricProcessor.ConsumeMetrics(context.Background(), metricServiceRequest)
		if err != nil {
			e.Telemetry.Logger.Error("Failed to consume metrics aggregated_net_request_gauge_group", zap.Error(err))
			return err
		}
		//TCP 链接指标
	case constnames.TcpStatsGaugeGroup:
		tcpInuseMetric := transform.GenerateXXMetric(gaugeGroup, constant.MetricTypeTcpStats)
		metricServiceRequest := transform.CreateFlattenMetrics(service, tcpInuseMetric)
		err := batchMetricProcessor.ConsumeMetrics(context.Background(), metricServiceRequest)
		if err != nil {
			e.Telemetry.Logger.Error("Failed to consume metrics tcp_metric_gauge_group", zap.Error(err))
			return err
		}
	case constnames.PageFaultGaugeGroupName:
		pageFaultMetric := transform.GenerateXXMetric(gaugeGroup, constant.MetricTypePageFault)
		metricServiceRequest := transform.CreateFlattenMetrics(service, pageFaultMetric)
		err := batchMetricProcessor.ConsumeMetrics(context.Background(), metricServiceRequest)
		if err != nil {
			e.Telemetry.Logger.Error("Failed to consume metrics page_fault_metric_gauge_group", zap.Error(err))
			return err
		}
	case constnames.ConnectGaugeGroupName:
		connectMetric := transform.GenerateConnectMetric(gaugeGroup, constant.MetricTypeConnect)
		metricServiceRequest := transform.CreateFlattenMetrics(service, connectMetric)
		err := batchMetricProcessor.ConsumeMetrics(context.Background(), metricServiceRequest)
		if err != nil {
			e.Telemetry.Logger.Error("Failed to consume metrics connect_metric_gauge_group", zap.Error(err))
			return err
		}
	default:
		return nil
	}
	return nil
}
