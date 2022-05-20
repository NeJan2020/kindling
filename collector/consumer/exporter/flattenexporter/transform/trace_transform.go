package transform

import (
	"github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/constant"
	flattenTraces "github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/data/protogen/collector/trace/v1"
	v11 "github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/data/protogen/common/v1"
	trace "github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/data/protogen/trace/v1"
	"github.com/Kindling-project/kindling/collector/model"
	"github.com/Kindling-project/kindling/collector/model/constlabels"
	"github.com/Kindling-project/kindling/collector/model/constvalues"
)

func CreateExportTraceServiceRequest(resourceSpans []*trace.ResourceSpans) flattenTraces.ExportTraceServiceRequest {
	initTraceRequest := flattenTraces.ExportTraceServiceRequest{
		ResourceSpans: resourceSpans,
	}
	return initTraceRequest
}

func GenerateResourceSpans(gaugeGroup *model.GaugeGroup) []*trace.ResourceSpans {
	return []*trace.ResourceSpans{{
		InstrumentationLibrarySpans: GenerateInstrumentationLibrarySpans(gaugeGroup),
	}}
}
func GenerateInstrumentationLibrarySpans(gaugeGroup *model.GaugeGroup) []*trace.InstrumentationLibrarySpans {
	return []*trace.InstrumentationLibrarySpans{{
		Spans: GenerateSpans(gaugeGroup),
	}}
}
func GenerateSpans(gaugeGroup *model.GaugeGroup) []*trace.Span {
	return []*trace.Span{{
		Events: GenerateEvents(gaugeGroup),
	}}
}
func GenerateEvents(gaugeGroup *model.GaugeGroup) []*trace.Span_Event {
	timestamp := gaugeGroup.Timestamp
	return []*trace.Span_Event{{
		TimeUnixNano: timestamp,
		Attributes:   GenerateAttributes(gaugeGroup),
	}}
}

func GenerateAttributes(gaugeGroup *model.GaugeGroup) []v11.KeyValue {
	keyValueSlice := make([]v11.KeyValue, 0)
	for _, gauge := range gaugeGroup.Values {
		GenerateKeyValueIntSlice(gauge.Name, gauge.GetInt().Value, &keyValueSlice)
	}
	labelMap := gaugeGroup.Labels

	//TODO gaugeGroup还没加这个字段
	GenerateKeyValueBoolSlice(constant.IsConnectFail, labelMap.GetBoolValue(constant.IsConnectFail), &keyValueSlice)
	GenerateKeyValueBoolSlice(constant.IsError, labelMap.GetBoolValue(constlabels.IsError), &keyValueSlice)
	GenerateKeyValueBoolSlice(constant.IsSlow, labelMap.GetBoolValue(constlabels.IsSlow), &keyValueSlice)
	GenerateKeyValueBoolSlice(constant.IsServer, labelMap.GetBoolValue(constlabels.IsServer), &keyValueSlice)

	GenerateKeyValueIntSlice(constant.Timestamp, labelMap.GetIntValue(constlabels.Timestamp), &keyValueSlice)
	GenerateKeyValueIntSlice(constant.Pid, labelMap.GetIntValue(constlabels.Pid), &keyValueSlice)
	GenerateKeyValueIntSlice(constant.Status, labelMap.GetIntValue(constlabels.HttpStatusCode), &keyValueSlice)
	GenerateKeyValueIntSlice(constant.SrcPort, labelMap.GetIntValue(constlabels.SrcPort), &keyValueSlice)
	GenerateKeyValueIntSlice(constant.DstPort, labelMap.GetIntValue(constlabels.DstPort), &keyValueSlice)

	protocol := labelMap.GetStringValue(constlabels.Protocol)
	GenerateKeyValueStringSlice(constant.APPProtocol, protocol, &keyValueSlice)
	GenerateKeyValueStringSlice(constant.SrcIp, labelMap.GetStringValue(constlabels.SrcIp), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.SrcNode, labelMap.GetStringValue(constlabels.SrcNode), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.SrcNamespace, labelMap.GetStringValue(constlabels.SrcNamespace), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.SrcWorkloadKind, labelMap.GetStringValue(constlabels.SrcWorkloadKind), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.SrcWorkloadName, labelMap.GetStringValue(constlabels.SrcWorkloadName), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.SrcPod, labelMap.GetStringValue(constlabels.SrcPod), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.SrcPodIp, labelMap.GetStringValue(constlabels.SrcIp), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.DstIp, labelMap.GetStringValue(constlabels.DstIp), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.DstPod, labelMap.GetStringValue(constlabels.DstPod), &keyValueSlice)

	DnatIp := labelMap.GetStringValue(constlabels.DnatIp)
	if DnatIp != "" {
		GenerateKeyValueStringSlice(constant.DstPodIp, DnatIp, &keyValueSlice)
	} else {
		GenerateKeyValueStringSlice(constant.DstPodIp, labelMap.GetStringValue(constlabels.DstIp), &keyValueSlice)
	}

	DnatPort := labelMap.GetStringValue(constlabels.DnatPort)
	if DnatPort != "" {
		GenerateKeyValueStringSlice(constant.DstPodPort, DnatPort, &keyValueSlice)
	} else {
		GenerateKeyValueStringSlice(constant.DstPodPort, labelMap.GetStringValue(constlabels.DstPort), &keyValueSlice)
	}

	GenerateKeyValueStringSlice(constant.DstContainerId, labelMap.GetStringValue(constlabels.DstContainerId), &keyValueSlice)

	protocolKey := constlabels.ContentKey
	/*if protocol == constvalues.ProtocolHttp || protocol == constvalues.ProtocolMysql {
		protocolKey = constlabels.ContentKey
	}*/
	if protocol == constvalues.ProtocolDns {
		protocolKey = constlabels.DnsDomain
	}
	if protocol == constvalues.ProtocolKafka {
		protocolKey = constlabels.KafkaTopic
	}
	GenerateKeyValueStringSlice(constant.ContentKey, labelMap.GetStringValue(protocolKey), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.ContainerId, labelMap.GetStringValue(constlabels.ContainerId), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.DstContainer, labelMap.GetStringValue(constlabels.DstContainer), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.DstNode, labelMap.GetStringValue(constlabels.DstNode), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.DstNamespace, labelMap.GetStringValue(constlabels.DstNamespace), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.DstWorkloadKind, labelMap.GetStringValue(constlabels.DstWorkloadKind), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.DstWorkloadName, labelMap.GetStringValue(constlabels.DstWorkloadName), &keyValueSlice)
	return keyValueSlice
}
