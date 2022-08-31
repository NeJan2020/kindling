package flattenexporter

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/constant"
	"github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/transform"
	"github.com/Kindling-project/kindling/collector/pkg/model"
	"github.com/Kindling-project/kindling/collector/pkg/model/constnames"
)

func (e *Cfg) Consume(dataGroup *model.DataGroup) error {
	if dataGroup == nil {
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
	if ce := e.Telemetry.Logger.Check(zap.DebugLevel, ""); ce != nil {
		e.Telemetry.Logger.Debug("exporter receives a dataGroup: \n" + dataGroup.String())
	}
	service := e.Config.GetServiceInstance()
	var err error
	switch dataGroup.Name {
	case constnames.SingleNetRequestMetricGroup:
		singleTrace := transform.GenerateResourceSpans(dataGroup)
		traceServiceRequest := transform.CreateExportTraceServiceRequest(singleTrace)
		err = batchTraceProcessor.ConsumeTraces(context.Background(), traceServiceRequest)
	case constnames.AggregatedNetRequestMetricGroup:
		requestMetric := transform.GenerateRequestMetric(dataGroup)
		metricServiceRequest := transform.CreateFlattenMetrics(service, requestMetric)
		err = batchMetricProcessor.ConsumeMetrics(context.Background(), metricServiceRequest)
		//TCP 链接指标
	case constnames.TcpStatsMetricGroup:
		tcpInuseMetric := transform.GenerateXXMetric(dataGroup, constant.MetricTypeTcpStats)
		metricServiceRequest := transform.CreateFlattenMetrics(service, tcpInuseMetric)
		err = batchMetricProcessor.ConsumeMetrics(context.Background(), metricServiceRequest)
	case constnames.PgftMetricGroupName:
		pageFaultMetric := transform.GenerateXXMetric(dataGroup, constant.MetricTypePageFault)
		metricServiceRequest := transform.CreateFlattenMetrics(service, pageFaultMetric)
		err = batchMetricProcessor.ConsumeMetrics(context.Background(), metricServiceRequest)
	case constnames.TcpConnectMetricGroupName:
		connectMetric := transform.GenerateConnectMetric(dataGroup, constant.MetricTypeConnect)
		metricServiceRequest := transform.CreateFlattenMetrics(service, connectMetric)
		err = batchMetricProcessor.ConsumeMetrics(context.Background(), metricServiceRequest)
	case constnames.TcpMetricGroupName:
	case constnames.ErrorSlowSyscallGroupName:
		syscallMetric := transform.GenerateXXMetric(dataGroup, constant.MetricTypeSysCall)
		metricServiceRequest := transform.CreateFlattenMetrics(service, syscallMetric)
		err = batchMetricProcessor.ConsumeMetrics(context.Background(), metricServiceRequest)
	case constnames.TcpStatusMetricGroup:
		tcpStatusMetric := transform.GenerateXXMetric(dataGroup, constant.MetricTypeTcpStatus)
		metricServiceRequest := transform.CreateFlattenMetrics(service, tcpStatusMetric)
		err = batchMetricProcessor.ConsumeMetrics(context.Background(), metricServiceRequest)
	case constnames.TcpSynAcceptQueueMetricGroupName:
		tcpBacklogMetric := transform.GenerateXXMetric(dataGroup, constant.MetricTypeTcpBacklog)
		metricServiceRequest := transform.CreateFlattenMetrics(service, tcpBacklogMetric)
		err = batchMetricProcessor.ConsumeMetrics(context.Background(), metricServiceRequest)
	default:
		err = fmt.Errorf("flatten exporter can't support to export this DataGroup: %s", dataGroup.Name)
	}
	if err != nil {
		e.Telemetry.Logger.Error("Failed to consume dataGroups", zap.String("DataGroupName", dataGroup.Name), zap.Error(err))
	}
	return err
}
