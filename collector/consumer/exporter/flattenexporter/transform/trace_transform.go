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

func GenerateResourceSpans(gaugeGroup *model.DataGroup) []*trace.ResourceSpans {
	return []*trace.ResourceSpans{{
		InstrumentationLibrarySpans: GenerateInstrumentationLibrarySpans(gaugeGroup),
	}}
}
func GenerateInstrumentationLibrarySpans(gaugeGroup *model.DataGroup) []*trace.InstrumentationLibrarySpans {
	return []*trace.InstrumentationLibrarySpans{{
		Spans: GenerateSpans(gaugeGroup),
	}}
}
func GenerateSpans(gaugeGroup *model.DataGroup) []*trace.Span {
	return []*trace.Span{{
		Events: GenerateEvents(gaugeGroup),
	}}
}
func GenerateEvents(gaugeGroup *model.DataGroup) []*trace.Span_Event {
	timestamp := gaugeGroup.Timestamp
	return []*trace.Span_Event{{
		TimeUnixNano: timestamp,
		Attributes:   GenerateAttributes(gaugeGroup),
	}}
}

func GenerateAttributes(gaugeGroup *model.DataGroup) []v11.KeyValue {
	keyValueSlice := make([]v11.KeyValue, 0, 50)
	for _, gauge := range gaugeGroup.Metrics {
		GenerateKeyValueIntSlice(gauge.Name, gauge.GetInt().Value, &keyValueSlice)
		//{Name: connect_time, Value: 0}
		//{Name: request_sent_time, Value: 9517}
		//{Name: waiting_ttfb_time, Value: 499799900}
		//{Name: content_download_time, Value: 79743}
		//{Name: request_total_time, Value: 499889160}
		//{Name: request_io, Value: 71}
		//{Name: response_io, Value: 22}
	}
	labelMap := gaugeGroup.Labels
	isServer := labelMap.GetBoolValue(constlabels.IsServer)
	GenerateKeyValueBoolSlice(constant.IsServer, isServer, &keyValueSlice)
	GenerateKeyValueBoolSlice(constant.IsError, labelMap.GetBoolValue(constlabels.IsError), &keyValueSlice)
	GenerateKeyValueBoolSlice(constant.IsSlow, labelMap.GetBoolValue(constlabels.IsSlow), &keyValueSlice)
	//GenerateKeyValueBoolSlice(constant.IsConnectFail, labelMap.GetBoolValue(constlabels.IsConnectFail), &keyValueSlice)
	protocol := labelMap.GetStringValue(constlabels.Protocol)
	GenerateKeyValueStringSlice(constant.APPProtocol, protocol, &keyValueSlice)

	protocolKey := constlabels.ContentKey
	var statusCode string
	if protocol == constvalues.ProtocolHttp {
		statusCode = constlabels.HttpStatusCode
		GenerateTraceApp(constant.RequestAPP, labelMap, &keyValueSlice)
		GenerateTraceApp(constant.ResponseAPP, labelMap, &keyValueSlice)
	}

	if protocol == constvalues.ProtocolDns {
		statusCode = constlabels.DnsRcode
		protocolKey = constlabels.DnsDomain
		//GenerateKeyValueIntSlice(constant.DnsQueryTime, labelMap.GetIntValue(constlabels.DnsQueryTime), &keyValueSlice)
		dnsRcode := labelMap.GetIntValue(constlabels.DnsRcode)
		GenerateKeyValueIntSlice(constant.DNS_R_CODE, dnsRcode, &keyValueSlice)
		GenerateKeyValueStringSlice(constant.DNS_DOMAIN, labelMap.GetStringValue(constlabels.DnsDomain), &keyValueSlice)
	}

	if protocol == constvalues.ProtocolKafka {
		statusCode = constlabels.KafkaErrorCode
		protocolKey = constlabels.KafkaTopic
	}

	if protocol == constvalues.ProtocolDubbo {
		statusCode = constlabels.DubboErrorCode
	}

	if protocol == constvalues.ProtocolMysql {
		protocolKey = constlabels.ContentKey
		statusCode = constlabels.SqlErrCode
	}

	if statusCode != "" {
		GenerateKeyValueIntSlice(constant.Status, labelMap.GetIntValue(statusCode), &keyValueSlice)
	}
	GenerateKeyValueStringSlice(constant.ContentKey, labelMap.GetStringValue(protocolKey), &keyValueSlice)
	GenerateKeyValueIntSlice(constant.Timestamp, int64(gaugeGroup.Timestamp), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.SrcIp, labelMap.GetStringValue(constlabels.SrcIp), &keyValueSlice)
	GenerateKeyValueIntSlice(constant.SrcPort, labelMap.GetIntValue(constlabels.SrcPort), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.SrcNode, labelMap.GetStringValue(constlabels.SrcNode), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.SrcNamespace, labelMap.GetStringValue(constlabels.SrcNamespace), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.SrcWorkloadKind, labelMap.GetStringValue(constlabels.SrcWorkloadKind), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.SrcWorkloadName, labelMap.GetStringValue(constlabels.SrcWorkloadName), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.SrcService, labelMap.GetStringValue(constlabels.SrcService), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.SrcPod, labelMap.GetStringValue(constlabels.SrcPod), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.SrcPodIp, labelMap.GetStringValue(constlabels.SrcIp), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.DstIp, labelMap.GetStringValue(constlabels.DstIp), &keyValueSlice)
	GenerateKeyValueIntSlice(constant.DstPort, labelMap.GetIntValue(constlabels.DstPort), &keyValueSlice)
	dnatIp := labelMap.GetStringValue(constlabels.DnatIp)
	if dnatIp != "" {
		GenerateKeyValueStringSlice(constant.DNatIp, dnatIp, &keyValueSlice)
		GenerateKeyValueStringSlice(constant.DstServiceIp, labelMap.GetStringValue(constlabels.DstIp), &keyValueSlice)
		GenerateKeyValueStringSlice(constant.DstPodIp, dnatIp, &keyValueSlice)
	} else {
		GenerateKeyValueStringSlice(constant.DstPodIp, labelMap.GetStringValue(constlabels.DstIp), &keyValueSlice)
	}
	dnatPort := labelMap.GetIntValue(constlabels.DnatPort)

	if dnatPort != -1 && dnatPort != 0 {
		GenerateKeyValueIntSlice(constant.DNatPort, dnatPort, &keyValueSlice)
		GenerateKeyValueIntSlice(constant.DstPodPort, dnatPort, &keyValueSlice)
		GenerateKeyValueIntSlice(constant.DstServicePort, labelMap.GetIntValue(constlabels.DstPort), &keyValueSlice)
	} else {
		GenerateKeyValueIntSlice(constant.DstPodPort, labelMap.GetIntValue(constlabels.DstPort), &keyValueSlice)
	}
	GenerateKeyValueStringSlice(constant.DstNode, labelMap.GetStringValue(constlabels.DstNode), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.DstNamespace, labelMap.GetStringValue(constlabels.DstNamespace), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.DstWorkloadKind, labelMap.GetStringValue(constlabels.DstWorkloadKind), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.DstWorkloadName, labelMap.GetStringValue(constlabels.DstWorkloadName), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.DstService, labelMap.GetStringValue(constlabels.DstService), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.DstPod, labelMap.GetStringValue(constlabels.DstPod), &keyValueSlice)
	GenerateKeyValueIntSlice(constant.Pid, labelMap.GetIntValue(constlabels.Pid), &keyValueSlice)
	var containerId string
	var containerName string
	if isServer {
		containerId = labelMap.GetStringValue(constlabels.DstContainerId)
		containerName = labelMap.GetStringValue(constlabels.DstContainer)
	} else {
		containerId = labelMap.GetStringValue(constlabels.SrcContainerId)
		containerName = labelMap.GetStringValue(constlabels.SrcContainer)
	}
	GenerateKeyValueStringSlice(constant.ContainerId, containerId, &keyValueSlice)
	GenerateKeyValueStringSlice(constant.ContainerName, containerName, &keyValueSlice)
	//GenerateKeyValueIntSlice(constant.HTTPS_TLS, labelMap.GetIntValue(constlabels.HTTPS_TLS), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.RequestPayload, labelMap.GetStringValue(constlabels.HttpRequestPayload), &keyValueSlice)
	GenerateKeyValueStringSlice(constant.ResponsePayload, labelMap.GetStringValue(constlabels.HttpResponsePayload), &keyValueSlice)
	//GenerateArrayValueSlice(constant.MESSAGE_CAPTURE, labelMap.GetStringValue(constlabels.HttpResponsePayload), &keyValueSlice)
	return keyValueSlice
}
