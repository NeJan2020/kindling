package transform

import (
	"strconv"

	"github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/constant"
	v1 "github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/data/protogen/common/v1"
	flattenMetrics "github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/data/protogen/metrics/flatten"
	"github.com/Kindling-project/kindling/collector/model"
	"github.com/Kindling-project/kindling/collector/model/constlabels"
	"github.com/Kindling-project/kindling/collector/model/constvalues"
)

func CreateFlattenMetrics(service *v1.Service, requestMetricArr []*flattenMetrics.RequestMetric) flattenMetrics.FlattenMetrics {
	requestMetrics := flattenMetrics.RequestMetrics{
		Service: service,
		Metrics: requestMetricArr,
	}
	initMetricRequest := flattenMetrics.FlattenMetrics{
		RequestMetricByte: &requestMetrics,
	}
	return initMetricRequest
}

func GenerateRequestMetric(gaugeGroup *model.DataGroup) []*flattenMetrics.RequestMetric {
	return []*flattenMetrics.RequestMetric{{
		MetricType:        constant.MetricTypeRequest,
		StartTimeUnixNano: gaugeGroup.Timestamp,
		MetricMap:         generateRequestMetricMap(gaugeGroup),
		Labels:            generateRequestMetricLabels(gaugeGroup),
	},
	}
}

func GenerateXXMetric(gaugeGroup *model.DataGroup, metricType int32) []*flattenMetrics.RequestMetric {
	return []*flattenMetrics.RequestMetric{{
		MetricType:        metricType,
		StartTimeUnixNano: gaugeGroup.Timestamp,
		MetricMap:         GenerateMetricMap(gaugeGroup),
		Labels:            GenerateMetricLabels(gaugeGroup),
	},
	}
}

func generateRequestMetricMap(gaugeGroup *model.DataGroup) map[string]*flattenMetrics.Metric {
	MetricMap := make(map[string]*flattenMetrics.Metric)
	gaugeMap := make(map[string]*model.Metric)
	for _, gauge := range gaugeGroup.Metrics {
		gaugeMap[gauge.Name] = gauge
	}
	MetricMap[constant.RequestIo] = generateMetric(constant.RequestIo, gaugeGroup, gaugeMap)
	MetricMap[constant.ResponseIo] = generateMetric(constant.ResponseIo, gaugeGroup, gaugeMap)
	MetricMap[constant.RequestDurationTime] = generateMetric(constant.RequestTotalTime, gaugeGroup, gaugeMap)
	MetricMap[constant.Error] = generateMetric(constant.Error, gaugeGroup, gaugeMap)
	MetricMap[constant.Slow] = generateMetric(constant.Slow, gaugeGroup, gaugeMap)
	MetricMap[constant.StatusCode1xxTotal] = generateMetric(constant.StatusCode1xxTotal, gaugeGroup, gaugeMap)
	MetricMap[constant.StatusCode2xxTotal] = generateMetric(constant.StatusCode2xxTotal, gaugeGroup, gaugeMap)
	MetricMap[constant.StatusCode3xxTotal] = generateMetric(constant.StatusCode3xxTotal, gaugeGroup, gaugeMap)
	MetricMap[constant.StatusCode4xxTotal] = generateMetric(constant.StatusCode4xxTotal, gaugeGroup, gaugeMap)
	MetricMap[constant.StatusCode5xxTotal] = generateMetric(constant.StatusCode5xxTotal, gaugeGroup, gaugeMap)
	return MetricMap
}

func GenerateConnectMetric(gaugeGroup *model.DataGroup, metricType int32) []*flattenMetrics.RequestMetric {
	return []*flattenMetrics.RequestMetric{{
		MetricType:        metricType,
		StartTimeUnixNano: gaugeGroup.Timestamp,
		MetricMap:         GenerateConnectMetricMap(gaugeGroup),
		Labels:            GenerateConnectMetricLabels(gaugeGroup),
	},
	}
}
func GenerateMetricMap(gaugeGroup *model.DataGroup) map[string]*flattenMetrics.Metric {
	metricMap := make(map[string]*flattenMetrics.Metric)
	for _, gauge := range gaugeGroup.Metrics {
		if gauge.DataType() == model.HistogramMetricType {
			metricMap[gauge.Name] = generateHistogramMetric(gauge.Name, gauge)
		}
		if gauge.DataType() == model.IntMetricType {
			metricMap[gauge.Name] = generateSumMetric(gauge.Name, gauge.GetInt().Value)
		}
	}
	return metricMap
}

func GenerateConnectMetricMap(gaugeGroup *model.DataGroup) map[string]*flattenMetrics.Metric {
	metricMap := make(map[string]*flattenMetrics.Metric)
	labelMap := gaugeGroup.Labels
	for _, gauge := range gaugeGroup.Metrics {
		if gauge.Name == constlabels.KindlingTcpConnectTotal {
			if gauge.DataType() == model.IntMetricType {
				if !labelMap.GetBoolValue(constlabels.Success) {
					metricMap[constant.ConnectFail] = generateSumMetric(constant.ConnectFail, gauge.GetInt().Value)
				}
			}
		}
		if gauge.Name == constlabels.KindlingTcpConnectDurationNanoseconds {
			if gauge.DataType() == model.HistogramMetricType {
				metricMap[constant.ConnectTime] = generateHistogramMetric(constant.ConnectTime, gauge)
			}
		}
	}
	return metricMap
}

func generateMetric(key string, gaugeGroup *model.DataGroup, gaugeMap map[string]*model.Metric) *flattenMetrics.Metric {
	metric, ok := gaugeMap[key]
	if ok {
		switch key {
		case constant.RequestIo, constant.ResponseIo:
			return generateSumMetric(key, metric.GetInt().Value)
		case constant.RequestTotalTime:
			if metric.DataType() == model.HistogramMetricType {
				return generateHistogramMetric(constant.RequestDurationTime, metric)
			} else {
				return generateSumMetric(key, metric.GetInt().Value)
			}
		case constant.Slow:
			isSlow := gaugeGroup.Labels.GetBoolValue(constlabels.IsSlow)
			if isSlow {
				return generateRequestCountMetric(key, gaugeMap)
			} else {
				return generateSumMetric(key, 0)
			}

		case constant.Error:
			isError := gaugeGroup.Labels.GetBoolValue(constlabels.IsError)
			if isError {
				return generateRequestCountMetric(key, gaugeMap)
			} else {
				return generateSumMetric(key, 0)
			}

		case constant.StatusCode1xxTotal:
			httpCode := gaugeGroup.Labels.GetIntValue(constlabels.HttpStatusCode)
			if httpCode < 200 {
				return generateRequestCountMetric(key, gaugeMap)
			} else {
				return generateSumMetric(key, 0)
			}
		case constant.StatusCode2xxTotal:
			httpCode := gaugeGroup.Labels.GetIntValue(constlabels.HttpStatusCode)
			if httpCode < 300 && httpCode >= 200 {
				return generateRequestCountMetric(key, gaugeMap)
			} else {
				return generateSumMetric(key, 0)
			}
		case constant.StatusCode3xxTotal:
			httpCode := gaugeGroup.Labels.GetIntValue(constlabels.HttpStatusCode)
			if httpCode < 400 && httpCode >= 300 {
				return generateRequestCountMetric(key, gaugeMap)
			} else {
				return generateSumMetric(key, 0)
			}
		case constant.StatusCode4xxTotal:
			httpCode := gaugeGroup.Labels.GetIntValue(constlabels.HttpStatusCode)
			if httpCode < 500 && httpCode >= 400 {
				return generateRequestCountMetric(key, gaugeMap)
			} else {
				return generateSumMetric(key, 0)
			}
		case constant.StatusCode5xxTotal:
			httpCode := gaugeGroup.Labels.GetIntValue(constlabels.HttpStatusCode)
			if httpCode >= 500 {
				return generateRequestCountMetric(key, gaugeMap)
			} else {
				return generateSumMetric(key, 0)
			}

		default:
			break
		}
	}
	return nil
}

func generateRequestCountMetric(key string, gaugeMap map[string]*model.Metric) *flattenMetrics.Metric {
	if gaugeMap[constant.RequestTotalTime].DataType() == model.HistogramMetricType {
		return generateSumMetric(key, int64(gaugeMap[constant.RequestTotalTime].GetHistogram().Count))
	} else {
		return generateSumMetric(key, gaugeMap[constvalues.RequestCount].GetInt().Value)
	}
}

func generateHistogramMetric(key string, gauge *model.Metric) *flattenMetrics.Metric {
	bucketCountsSlice := make([]float64, len(gauge.GetHistogram().ExplicitBoundaries))
	bucketCountsFloatSlice := gauge.GetHistogram().ExplicitBoundaries
	for i, value := range bucketCountsFloatSlice {
		bucketCountsSlice[i] = float64(value)
	}
	return &flattenMetrics.Metric{Name: key, Data: &flattenMetrics.Metric_Histogram{Histogram: &flattenMetrics.Histogram{
		Count:          gauge.GetHistogram().Count,
		Sum:            gauge.GetHistogram().Sum,
		BucketCounts:   gauge.GetHistogram().BucketCounts,
		ExplicitBounds: bucketCountsSlice,
	}}}
}

func generateSumMetric(key string, value int64) *flattenMetrics.Metric {
	return &flattenMetrics.Metric{Name: key, Data: &flattenMetrics.Metric_Sum{Sum: &flattenMetrics.Sum{Value: value}}}
}

func generateRequestMetricLabels(gaugeGroup *model.DataGroup) []v1.StringKeyValue {
	metricLabels := make([]v1.StringKeyValue, 0, 27)
	labelMap := gaugeGroup.Labels
	generateK8sLabels(labelMap, metricLabels)
	protocol := labelMap.GetStringValue(constlabels.Protocol)
	GenerateStringKeyValueSlice(constant.Protocol, protocol, &metricLabels)
	protocolKey := constlabels.ContentKey
	if protocol == constvalues.ProtocolHttp {
	}

	if protocol == constvalues.ProtocolDns {
		protocolKey = constlabels.DnsDomain
	}

	if protocol == constvalues.ProtocolKafka {
		protocolKey = constlabels.KafkaTopic
	}

	if protocol == constvalues.ProtocolDubbo {
	}

	if protocol == constvalues.ProtocolMysql {
		protocolKey = constlabels.KafkaTopic
	}
	GenerateStringKeyValueSlice(constant.ContentKey, labelMap.GetStringValue(protocolKey), &metricLabels)
	return metricLabels
}

func generateK8sLabels(labelMap *model.AttributeMap, metricLabels []v1.StringKeyValue) {
	GenerateStringKeyValueSlice(constant.Pid, strconv.FormatInt(labelMap.GetIntValue(constlabels.Pid), 10), &metricLabels)
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
	dnatIp := labelMap.GetStringValue(constlabels.DnatIp)
	GenerateStringKeyValueSlice(constant.DstIp, labelMap.GetStringValue(constlabels.DstIp), &metricLabels)
	GenerateStringKeyValueSlice(constant.DNatIp, dnatIp, &metricLabels)
	GenerateStringKeyValueSlice(constant.DstContainer, labelMap.GetStringValue(constlabels.DstContainer), &metricLabels)
	GenerateStringKeyValueSlice(constant.DstContainerId, labelMap.GetStringValue(constlabels.DstContainerId), &metricLabels)
	GenerateStringKeyValueSlice(constant.IsServer, strconv.FormatBool(labelMap.GetBoolValue(constlabels.IsServer)), &metricLabels)
	var dstServiceIp = ""
	var dstServicePort = ""
	if dnatIp != "" {
		dstServiceIp = labelMap.GetStringValue(constlabels.DstIp)
		GenerateStringKeyValueSlice(constant.DstPodIp, dnatIp, &metricLabels)
	} else {
		GenerateStringKeyValueSlice(constant.DstPodIp, labelMap.GetStringValue(constlabels.DstIp), &metricLabels)
	}
	GenerateStringKeyValueSlice(constant.DstServiceIp, dstServiceIp, &metricLabels)
	dnatPort := labelMap.GetIntValue(constlabels.DnatPort)
	GenerateStringKeyValueSlice(constant.DNatPort, strconv.FormatInt(dnatPort, 10), &metricLabels)
	if dnatPort != -1 && dnatPort != 0 {
		dstServicePort = strconv.FormatInt(dnatPort, 10)
		GenerateStringKeyValueSlice(constant.DstPodPort, strconv.FormatInt(dnatPort, 10), &metricLabels)
	} else {
		GenerateStringKeyValueSlice(constant.DstPodPort, strconv.FormatInt(labelMap.GetIntValue(constlabels.DstPort), 10), &metricLabels)
	}
	GenerateStringKeyValueSlice(constant.DstServicePort, dstServicePort, &metricLabels)
}

func GenerateConnectMetricLabels(gaugeGroup *model.DataGroup) []v1.StringKeyValue {
	metricLabels := make([]v1.StringKeyValue, 0, 25)
	labelsMap := gaugeGroup.Labels
	generateK8sLabels(labelsMap, metricLabels)
	return metricLabels
}

func GenerateMetricLabels(gaugeGroup *model.DataGroup) []v1.StringKeyValue {
	metricLabels := make([]v1.StringKeyValue, 0, 25)
	labelsMap := gaugeGroup.Labels.ToStringMap()
	for k, v := range labelsMap {
		GenerateStringKeyValueSlice(k, v, &metricLabels)
	}
	return metricLabels
}
