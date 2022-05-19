package transform

import (
	"github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/constant"
	v1 "github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/data/protogen/common/v1"
	flattenMetrics "github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/data/protogen/metrics/flatten"
	"github.com/Kindling-project/kindling/collector/model"
	"github.com/Kindling-project/kindling/collector/model/constlabels"
	"github.com/Kindling-project/kindling/collector/model/constvalues"
	"strconv"
)

func CreateFlattenMetrics(service *v1.Service, requestMetricArr []*flattenMetrics.RequestMetric) flattenMetrics.FlattenMetrics {
	requestMetrics := flattenMetrics.RequestMetrics{
		Service: service,
		Metrics: requestMetricArr,
	}
	initMetricRequest := flattenMetrics.FlattenMetrics{
		RequestMetrics: &requestMetrics,
	}
	return initMetricRequest
}

func GenerateRequestMetrics(gaugeGroup *model.GaugeGroup, service *v1.Service) *flattenMetrics.RequestMetrics {
	return &flattenMetrics.RequestMetrics{
		Service: service,
		Metrics: GenerateRequestMetric(gaugeGroup),
	}
}

func GenerateRequestMetric(gaugeGroup *model.GaugeGroup) []*flattenMetrics.RequestMetric {
	return []*flattenMetrics.RequestMetric{{
		MetricType:        constant.MetricTypeRequest,
		StartTimeUnixNano: gaugeGroup.Timestamp,
		MetricMap:         GenerateRequestMetricMap(gaugeGroup),
		Labels:            GenerateRequestMetricLabels(gaugeGroup),
	},
	}
}

func GenerateTcpInuseMetric(gaugeGroup *model.GaugeGroup) []*flattenMetrics.RequestMetric {
	return []*flattenMetrics.RequestMetric{{
		MetricType:        constant.MetricTypeTcpInuse,
		StartTimeUnixNano: gaugeGroup.Timestamp,
		MetricMap:         GenerateTcpInuseMetricMap(gaugeGroup),
		Labels:            GenerateTcpInuseMetricLabels(gaugeGroup),
	},
	}
}

func GenerateRequestMetricMap(gaugeGroup *model.GaugeGroup) map[string]*flattenMetrics.Metric {
	MetricMap := make(map[string]*flattenMetrics.Metric)
	gaugeMap := make(map[string]*model.Gauge)
	for _, gauge := range gaugeGroup.Values {
		gaugeMap[gauge.Name] = gauge
	}
	MetricMap[constant.RequestIo] = GenerateMetric(constant.RequestIo, gaugeGroup, gaugeMap)
	MetricMap[constant.ResponseIo] = GenerateMetric(constant.ResponseIo, gaugeGroup, gaugeMap)
	MetricMap[constant.RequestDurationTime] = GenerateMetric(constant.RequestDurationTime, gaugeGroup, gaugeMap)
	MetricMap[constant.Error] = GenerateMetric(constant.Error, gaugeGroup, gaugeMap)
	MetricMap[constant.Slow] = GenerateMetric(constant.Slow, gaugeGroup, gaugeMap)
	MetricMap[constant.RequestTotalTime] = GenerateMetric(constant.RequestTotalTime, gaugeGroup, gaugeMap)
	MetricMap[constant.StatusCode1xxTotal] = GenerateMetric(constant.StatusCode1xxTotal, gaugeGroup, gaugeMap)
	MetricMap[constant.StatusCode2xxTotal] = GenerateMetric(constant.StatusCode2xxTotal, gaugeGroup, gaugeMap)
	MetricMap[constant.StatusCode3xxTotal] = GenerateMetric(constant.StatusCode3xxTotal, gaugeGroup, gaugeMap)
	MetricMap[constant.StatusCode4xxTotal] = GenerateMetric(constant.StatusCode4xxTotal, gaugeGroup, gaugeMap)
	MetricMap[constant.StatusCode5xxTotal] = GenerateMetric(constant.StatusCode5xxTotal, gaugeGroup, gaugeMap)
	return MetricMap
}

func GenerateTcpInuseMetricMap(gaugeGroup *model.GaugeGroup) map[string]*flattenMetrics.Metric {
	metricMap := make(map[string]*flattenMetrics.Metric)
	for _, gauge := range gaugeGroup.Values {
		if gauge.DataType() == model.IntGaugeType {
			metricMap[gauge.Name] = GenerateSumMetric(gauge.Name, gauge)
		}
	}
	return metricMap
}

func GenerateMetric(key string, gaugeGroup *model.GaugeGroup, gaugeMap map[string]*model.Gauge) *flattenMetrics.Metric {
	switch key {
	case constant.RequestIo:
		return GenerateSumMetric(key, gaugeMap[key])
	case constant.ResponseIo:
		return GenerateSumMetric(key, gaugeMap[key])
	case constant.RequestTotalTime:
		return GenerateSumMetric(key, gaugeMap[key])
	case constant.RequestDurationTime:
		return GenerateHistogramMetric(key, gaugeMap)
	case constant.Slow:
		if gaugeMap[constvalues.RequestTotalTime].GetInt().Value > 500 {
			return GenerateSumMetric(key, gaugeMap[constvalues.RequestCount])
		} else {
			return GenerateSumMetric(key, model.NewIntGauge(key, 0))
		}
	case constant.Error:
		if gaugeGroup.Labels.GetIntValue(constlabels.HttpStatusCode) > 400 {
			return GenerateSumMetric(key, gaugeMap[constvalues.RequestCount])
		} else {
			return GenerateSumMetric(key, model.NewIntGauge(key, 0))
		}
	case constant.StatusCode1xxTotal:
		httpCode := gaugeGroup.Labels.GetIntValue(constlabels.HttpStatusCode)
		if httpCode < 200 {
			return GenerateSumMetric(key, gaugeMap[constvalues.RequestCount])
		} else {
			return GenerateSumMetric(key, model.NewIntGauge(key, 0))
		}
	case constant.StatusCode2xxTotal:
		httpCode := gaugeGroup.Labels.GetIntValue(constlabels.HttpStatusCode)
		if httpCode < 300 && httpCode >= 200 {
			return GenerateSumMetric(key, gaugeMap[constvalues.RequestCount])
		} else {
			return GenerateSumMetric(key, model.NewIntGauge(key, 0))
		}
	case constant.StatusCode3xxTotal:
		httpCode := gaugeGroup.Labels.GetIntValue(constlabels.HttpStatusCode)
		if httpCode < 400 && httpCode >= 300 {
			return GenerateSumMetric(key, gaugeMap[constvalues.RequestCount])
		} else {
			return GenerateSumMetric(key, model.NewIntGauge(key, 0))
		}
	case constant.StatusCode4xxTotal:
		httpCode := gaugeGroup.Labels.GetIntValue(constlabels.HttpStatusCode)
		if httpCode < 500 && httpCode >= 400 {
			return GenerateSumMetric(key, gaugeMap[constvalues.RequestCount])
		} else {
			return GenerateSumMetric(key, model.NewIntGauge(key, 0))
		}
	case constant.StatusCode5xxTotal:
		httpCode := gaugeGroup.Labels.GetIntValue(constlabels.HttpStatusCode)
		if httpCode >= 500 {
			return GenerateSumMetric(key, gaugeMap[constvalues.RequestCount])
		} else {
			return GenerateSumMetric(key, model.NewIntGauge(key, 0))
		}
	default:
		break
	}
	return nil
}

func GenerateSumMetric(key string, value *model.Gauge) *flattenMetrics.Metric {
	return &flattenMetrics.Metric{Name: key, Data: &flattenMetrics.Metric_Sum{Sum: &flattenMetrics.Sum{Value: value.GetInt().Value}}}
}

func GenerateHistogramMetric(key string, gaugeMap map[string]*model.Gauge) *flattenMetrics.Metric {
	bucketCountsSlice := make([]float64, len(gaugeMap[key].GetHistogram().ExplicitBoundaries))
	bucketCountsFloatSlice := gaugeMap[key].GetHistogram().ExplicitBoundaries
	for i, value := range bucketCountsFloatSlice {
		bucketCountsSlice[i] = float64(value)
	}
	return &flattenMetrics.Metric{Name: key, Data: &flattenMetrics.Metric_Histogram{Histogram: &flattenMetrics.Histogram{
		Count:          gaugeMap[key].GetHistogram().Count,
		Sum:            gaugeMap[key].GetHistogram().Sum,
		BucketCounts:   gaugeMap[key].GetHistogram().BucketCounts,
		ExplicitBounds: bucketCountsSlice,
	}}}
}

func GenerateRequestMetricLabels(gaugeGroup *model.GaugeGroup) []v1.StringKeyValue {
	metricLabels := make([]v1.StringKeyValue, 0)
	labelMap := gaugeGroup.Labels
	GenerateStringKeyValueSlice(constant.Pid, strconv.FormatInt(labelMap.GetIntValue(constlabels.Pid), 10), &metricLabels)
	GenerateStringKeyValueSlice(constant.SrcMasterIP, "SrcMasterIP", &metricLabels)
	GenerateStringKeyValueSlice(constant.DstMasterIP, "DstMasterIP", &metricLabels)
	GenerateStringKeyValueSlice(constant.SrcNode, labelMap.GetStringValue(constlabels.SrcNode), &metricLabels)
	GenerateStringKeyValueSlice(constant.SrcNamespace, labelMap.GetStringValue(constlabels.SrcNamespace), &metricLabels)
	GenerateStringKeyValueSlice(constant.SrcWorkloadKind, labelMap.GetStringValue(constlabels.SrcWorkloadKind), &metricLabels)
	GenerateStringKeyValueSlice(constant.SrcWorkloadName, labelMap.GetStringValue(constlabels.SrcWorkloadName), &metricLabels)
	GenerateStringKeyValueSlice(constant.SrcService, labelMap.GetStringValue(constlabels.SrcService), &metricLabels)
	GenerateStringKeyValueSlice(constant.SrcPod, labelMap.GetStringValue(constlabels.SrcPod), &metricLabels)
	GenerateStringKeyValueSlice(constant.SrcContainer, labelMap.GetStringValue(constlabels.SrcContainer), &metricLabels)
	GenerateStringKeyValueSlice(constant.SrcContainerId, labelMap.GetStringValue(constlabels.SrcContainerId), &metricLabels)
	GenerateStringKeyValueSlice(constant.SrcIp, labelMap.GetStringValue(constlabels.SrcIp), &metricLabels)
	GenerateStringKeyValueSlice(constant.DstNode, labelMap.GetStringValue(constlabels.DstNode), &metricLabels)
	GenerateStringKeyValueSlice(constant.DstNamespace, labelMap.GetStringValue(constlabels.DstNamespace), &metricLabels)
	GenerateStringKeyValueSlice(constant.DstWorkloadKind, labelMap.GetStringValue(constlabels.DstWorkloadKind), &metricLabels)
	GenerateStringKeyValueSlice(constant.DstWorkloadName, labelMap.GetStringValue(constlabels.DstWorkloadName), &metricLabels)
	GenerateStringKeyValueSlice(constant.DstService, labelMap.GetStringValue(constlabels.DstService), &metricLabels)
	GenerateStringKeyValueSlice(constant.DstPod, labelMap.GetStringValue(constlabels.DstPod), &metricLabels)
	GenerateStringKeyValueSlice(constant.DstIp, labelMap.GetStringValue(constlabels.DstIp), &metricLabels)
	GenerateStringKeyValueSlice(constant.DnatIp, labelMap.GetStringValue(constlabels.DnatIp), &metricLabels)
	GenerateStringKeyValueSlice(constant.DstServiceIp, "DstServiceIp", &metricLabels)
	GenerateStringKeyValueSlice(constant.DstServicePort, "DstServicePort", &metricLabels)
	GenerateStringKeyValueSlice(constant.DstPodIp, "DstPodIp", &metricLabels)
	GenerateStringKeyValueSlice(constant.DstPodPort, "DstPodPort", &metricLabels)
	GenerateStringKeyValueSlice(constant.DstContainer, labelMap.GetStringValue(constlabels.DstContainer), &metricLabels)
	GenerateStringKeyValueSlice(constant.DstContainerId, labelMap.GetStringValue(constlabels.DstContainerId), &metricLabels)
	GenerateStringKeyValueSlice(constant.IsServer, strconv.FormatBool(labelMap.GetBoolValue(constlabels.IsServer)), &metricLabels)
	GenerateStringKeyValueSlice(constant.HttpHost, "HttpHost", &metricLabels)

	protocol := labelMap.GetStringValue(constlabels.Protocol)
	GenerateStringKeyValueSlice(constant.Protocol, labelMap.GetStringValue(protocol), &metricLabels)
	var protocolKey string
	if protocol == constvalues.ProtocolHttp || protocol == constvalues.ProtocolMysql {
		protocolKey = constlabels.ContentKey
	}
	if protocol == constvalues.ProtocolDns {
		protocolKey = constlabels.DnsDomain
	}
	if protocol == constvalues.ProtocolKafka {
		protocolKey = constlabels.KafkaTopic
	}
	GenerateStringKeyValueSlice(constant.ContentKey, labelMap.GetStringValue(protocolKey), &metricLabels)
	return metricLabels
}

func GenerateTcpInuseMetricLabels(gaugeGroup *model.GaugeGroup) []v1.StringKeyValue {
	metricLabels := make([]v1.StringKeyValue, 0)

	labelsMap := gaugeGroup.Labels.ToStringMap()
	for k, v := range labelsMap {
		GenerateStringKeyValueSlice(k, v, &metricLabels)
	}
	return metricLabels
}

func CreateDefaultExplicitBoundsSlice() []float64 {
	explicitBoundsSlice := make([]float64, 0, 1)
	//TODO 需要修改
	boundsArr := []float64{100, 200, 300, 500, 1000, 2000, 3000, 5000, 10000, 20000, 30000, 50000, 100000}
	explicitBoundsSlice = append(explicitBoundsSlice, boundsArr...)
	return explicitBoundsSlice
}

func CreateBucketCountsSlice(explicitBoundsSlice []float64, requestAvgTime int64, requestCount int64) []uint64 {
	bucketCountsSlice := make([]uint64, 0, 1)
	for _, bucket := range explicitBoundsSlice {
		if float64(requestAvgTime) > bucket {
			bucketCountsSlice = append(bucketCountsSlice, 0)
		} else {
			bucketCountsSlice = append(bucketCountsSlice, uint64(requestCount))
		}
	}
	return bucketCountsSlice
}
