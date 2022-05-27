package flattenexporter

import (
	"github.com/Kindling-project/kindling/collector/component"
	"github.com/Kindling-project/kindling/collector/model"
	"github.com/Kindling-project/kindling/collector/model/constlabels"
	"github.com/Kindling-project/kindling/collector/model/constnames"
	"github.com/Kindling-project/kindling/collector/model/constvalues"
	"github.com/spf13/viper"
	"strconv"
	"sync"
	"testing"
)

func makeSingleGaugeGroup(i int) *model.GaugeGroup {
	gaugesGroup := &model.GaugeGroup{
		Name: constnames.SingleNetRequestGaugeGroup,
		Values: []*model.Gauge{
			model.NewIntGauge(constvalues.ResponseIo, 123456789),
			model.NewIntGauge(constvalues.RequestIo, 1234567891),
			model.NewIntGauge(constvalues.RequestTotalTime, int64(i+9)),
			model.NewIntGauge(constvalues.ConnectTime, 4500),
			model.NewIntGauge(constvalues.RequestSentTime, 4500),
			model.NewIntGauge(constvalues.WaitingTtfbTime, 4500),
			model.NewIntGauge(constvalues.ContentDownloadTime, 4500),
			model.NewIntGauge(constvalues.RequestCount, 4500),
			//{Name: connect_time, Value: 0}
			//{Name: request_sent_time, Value: 9517}
			//{Name: waiting_ttfb_time, Value: 499799900}
			//{Name: content_download_time, Value: 79743}
			//{Name: request_total_time, Value: 499889160}
			//{Name: request_io, Value: 71}
			//{Name: response_io, Value: 22}
		},
		Labels:    model.NewAttributeMap(),
		Timestamp: 19900909090,
	}
	gaugesGroup.Labels.AddBoolValue(constlabels.IsSlow, true)
	gaugesGroup.Labels.AddBoolValue(constlabels.IsError, true)
	gaugesGroup.Labels.AddBoolValue(constlabels.IsServer, true)
	gaugesGroup.Labels.AddIntValue(constlabels.Timestamp, 19900909090)
	gaugesGroup.Labels.AddStringValue(constlabels.SrcIp, "test-SrcIp"+strconv.Itoa(i))
	gaugesGroup.Labels.AddIntValue(constlabels.SrcPort, 8080)
	gaugesGroup.Labels.AddStringValue(constlabels.SrcNode, "test-SrcNode"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.SrcNamespace, "test-SrcNamespace"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.SrcWorkloadKind, "test-SrcWorkloadKind"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.SrcWorkloadName, "test-SrcWorkloadName"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.SrcService, "test-SrcService"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.SrcPod, "test-SrcPod"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.DstIp, "test-dstIp"+strconv.Itoa(i))
	gaugesGroup.Labels.AddIntValue(constlabels.DstPort, 8090)
	gaugesGroup.Labels.AddStringValue(constlabels.DnatIp, "test-dnatIp"+strconv.Itoa(i))
	gaugesGroup.Labels.AddIntValue(constlabels.DnatPort, 8091)
	gaugesGroup.Labels.AddStringValue(constlabels.DstNode, "test-DstNode"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.DstNamespace, "test-DstNamespace"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.DstWorkloadName, "test-DstWorkloadName"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.DstWorkloadKind, "test-DstWorkloadKind"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.DstService, "test-DstService"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.DstPod, "test-DstPod"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.HttpMethod, "test-HttpMethod"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.HttpUrl, "test-HttpUrl"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.HttpRequestPayload, "GET / HTTP/1.1\\r\\nHost: :8080\\r\\nUser-Agent: Go-http-client/1.1\\r\\nAccept-Encoding: gz")
	gaugesGroup.Labels.AddStringValue(constlabels.HttpRequestPayload, "HTTP/1.1 404 Not Found\\r\\nContent-Type: text/plain; charset=utf-8\\r\\nX-Content-Type-")
	gaugesGroup.Labels.AddStringValue(constlabels.SrcContainerId, "test-SrcContainerId"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.SrcContainer, "test-SrcContainerName"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.DstContainerId, "test-DstContainerId"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.DstContainer, "test-DstContainerName"+strconv.Itoa(i))
	gaugesGroup.Labels.AddIntValue(constlabels.HttpStatusCode, 299)
	gaugesGroup.Labels.AddIntValue(constlabels.KafkaErrorCode, 298)
	gaugesGroup.Labels.AddIntValue(constlabels.SqlErrCode, 297)
	gaugesGroup.Labels.AddIntValue(constlabels.Pid, int64(i))
	gaugesGroup.Labels.AddStringValue(constlabels.Protocol, "dns")
	gaugesGroup.Labels.AddStringValue(constlabels.ContentKey, "/url")
	gaugesGroup.Labels.AddIntValue(constlabels.DnsRcode, int64(i+200))
	gaugesGroup.Labels.AddStringValue(constlabels.DnsDomain, "DNS_DOMAIN")
	return gaugesGroup
}

func makeAggNetGaugeGroup(i int) *model.GaugeGroup {
	gaugesGroup := &model.GaugeGroup{
		Name: constnames.AggregatedNetRequestGaugeGroup,
		Values: []*model.Gauge{
			model.NewHistogramGauge(constvalues.RequestTotalTime, &model.Histogram{
				Sum:                10000*4999 + 15000,
				Count:              10000,
				ExplicitBoundaries: []int64{0, 100, 200, 500, 1000, 2000, 5000, 10000},
				BucketCounts:       []uint64{0, 100, 200, 500, 1000, 2000, 5000, 10000},
			}),
			model.NewIntGauge(constvalues.ResponseIo, 1234567891),
			model.NewIntGauge(constvalues.RequestIo, 4500),
		},
		Labels:    model.NewAttributeMap(),
		Timestamp: 19900909090,
	}
	gaugesGroup.Labels.AddStringValue(constlabels.SrcNode, "test-SrcNode"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.SrcNamespace, "test-SrcNamespace"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.SrcPod, "test-SrcPod"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.SrcWorkloadName, "test-SrcWorkloadName"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.SrcWorkloadKind, "test-SrcWorkloadKind"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.SrcService, "test-SrcService"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.SrcIp, "test-SrcIp"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.DstNode, "test-DstNode"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.DstNamespace, "test-DstNamespace"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.DstPod, "test-DstPod"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.DstWorkloadName, "test-DstWorkloadName"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.DstWorkloadKind, "test-DstWorkloadKind"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.DstService, "test-DstService"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.SrcContainer, "test-SrcContainer"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.SrcContainerId, "test-SrcContainerId"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.DstContainer, "test-SrcContainer"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.DstContainerId, "test-SrcContainerId"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.Protocol, "kafka")
	//gaugesGroup.Labels.AddStringValue(constlabels.Protocol, "http")
	gaugesGroup.Labels.AddIntValue(constlabels.HttpStatusCode, 498)
	gaugesGroup.Labels.AddBoolValue(constlabels.IsSlow, true)
	gaugesGroup.Labels.AddBoolValue(constlabels.IsError, true)
	gaugesGroup.Labels.AddBoolValue(constlabels.IsServer, true)
	gaugesGroup.Labels.AddStringValue(constlabels.ContentKey, "/http/url")
	// Topology data preferentially use D Nat Ip and D Nat Port
	gaugesGroup.Labels.AddStringValue(constlabels.DstIp, "test-DstIp")
	gaugesGroup.Labels.AddIntValue(constlabels.DstPort, 8081)
	gaugesGroup.Labels.AddStringValue(constlabels.KafkaTopic, "npm_detail_topology_request")
	gaugesGroup.Labels.AddStringValue(constlabels.DnatIp, "test-DnatIp"+strconv.Itoa(i))
	gaugesGroup.Labels.AddIntValue(constlabels.DnatPort, 8082)
	gaugesGroup.Labels.AddIntValue(constlabels.Pid, 111)
	return gaugesGroup
}
func makeTcpStatsGaugeGroup(i int) *model.GaugeGroup {
	gaugesGroup := &model.GaugeGroup{
		Name: constnames.TcpStatsGaugeGroup,
		Values: []*model.Gauge{
			model.NewIntGauge("Established", int64(i)),
			model.NewIntGauge("SynSent", int64(i)),
			model.NewIntGauge("SynRecv", int64(i)),
			model.NewIntGauge("FinWait1", int64(i)),
			model.NewIntGauge("FinWait2", int64(i)),
			model.NewIntGauge("TimeWait", int64(i)),
			model.NewIntGauge("Close", int64(i)),
			model.NewIntGauge("CloseWait", int64(i)),
			model.NewIntGauge("LastAck", int64(i)),
			model.NewIntGauge("Listen", int64(i)),
			model.NewIntGauge("Closing", int64(i)),
		},
		Labels:    model.NewAttributeMap(),
		Timestamp: 19900909090,
	}
	gaugesGroup.Labels.AddStringValue("mode", "tcp4")
	gaugesGroup.Labels.AddStringValue("container", "test-container"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue("container_id", "test-container_id"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue("namespace", "test-namespace"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue("node", "test-node"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue("node_ip", "test-node_ip"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue("pod", "test-pod"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue("service", "test-elasticsearch-svc"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue("workload_kind", "test-statefulset"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue("workload_name", "test-workload_name"+strconv.Itoa(i))
	return gaugesGroup
}

func makePageFaultGaugeGroup(i int) *model.GaugeGroup {
	gaugesGroup := &model.GaugeGroup{
		Name: constnames.PageFaultGaugeGroupName,
		Values: []*model.Gauge{
			model.NewIntGauge("kindling_pagefault_major_total", int64(i)),
			model.NewIntGauge("kindling_pagefault_minor_total", int64(i)),
		},
		Labels:    model.NewAttributeMap(),
		Timestamp: 19900909090,
	}
	gaugesGroup.Labels.AddStringValue("tid", "test-tid"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue("pid", "test-pid"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue("container_id", "test-container_id"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue("namespace", "test-namespace"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue("node", "test-node"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue("node_ip", "test-node_ip"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue("pod", "test-pod"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue("service", "test-elasticsearch-svc"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue("workload_kind", "test-statefulset"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue("workload_name", "test-workload_name"+strconv.Itoa(i))
	return gaugesGroup
}

func makeTcpMetricGaugeGroup(i int) *model.GaugeGroup {
	gaugesGroup := &model.GaugeGroup{
		Name: constnames.ConnectGaugeGroupName,
		Values: []*model.Gauge{
			model.NewIntGauge(constlabels.KindlingTcpConnectTotal, int64(i)),
			model.NewHistogramGauge(constlabels.KindlingTcpConnectDurationNanoseconds, &model.Histogram{
				Sum:                15000,
				Count:              10000,
				ExplicitBoundaries: []int64{0, 100, 200, 500, 1000, 2000, 5000, 10000},
				BucketCounts:       []uint64{0, 100, 200, 500, 1000, 2000, 5000, 10000},
			}),
		},
		Labels:    model.NewAttributeMap(),
		Timestamp: 19900909090,
	}

	gaugesGroup.Labels.AddStringValue(constlabels.DstContainer, "test-DstContainer"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.DstContainerId, "test-DstContainerId"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.DstIp, "test-DnatIp")
	gaugesGroup.Labels.AddIntValue(constlabels.DstPort, 8081)
	gaugesGroup.Labels.AddStringValue(constlabels.DstNode, "test-DstNode"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.DstNodeIp, "test-DstNodeIp"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.DstNamespace, "test-DstNamespace"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.DstPod, "test-DstPod"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.DstWorkloadName, "test-DstWorkloadName"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.DstWorkloadKind, "test-DstWorkloadKind"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.DstService, "test-DstService"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.SrcContainer, "test-SrcContainer"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.SrcContainerId, "test-SrcContainerId"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.SrcIp, "test-SrcIp"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.SrcPort, "test-SrcPort"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.SrcNode, "test-SrcNode"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.SrcNodeIp, "test-SrcNodeIp"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.SrcNamespace, "test-SrcNamespace"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.SrcPod, "test-SrcPod"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.SrcWorkloadName, "test-SrcWorkloadName"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.SrcWorkloadKind, "test-SrcWorkloadKind"+strconv.Itoa(i))
	gaugesGroup.Labels.AddStringValue(constlabels.SrcService, "test-SrcService"+strconv.Itoa(i))
	gaugesGroup.Labels.AddBoolValue(constlabels.IsServer, false)
	gaugesGroup.Labels.AddBoolValue(constlabels.Success, false)
	return gaugesGroup
}

func TestInitFlattenExporter(t *testing.T) {
	InitFlattenExporter(t)
}

func InitFlattenExporter(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	configPath := "testdata/kindling-collector-config.yml"
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		t.Fatalf("error happened when reading config: %v", err)
	}
	config := &Config{}
	err = viper.UnmarshalKey("exporters.flattenexporter", config)
	if err != nil {
		t.Fatalf("error happened when unmarshaling config: %v", err)
	}
	export := NewExporter(config, component.NewDefaultTelemetryTools())
	for i := 0; i < 1; i++ {
		//go export.Consume(makeSingleGaugeGroup(i))
		//time.Sleep(1 * time.Second)
		//go export.Consume(makePageFaultGaugeGroup(i))
	}

	for i := 1; i < 2; i++ {
		//go export.Consume(makeSingleGaugeGroup(i))
		go export.Consume(makeAggNetGaugeGroup(i))
		//go export.Consume(makeTcpStatsGaugeGroup(i))
		//go export.Consume(makePageFaultGaugeGroup(i))

		//go export.Consume(makeTcpMetricGaugeGroup(i))
	}
	for i := 0; i < 10; i++ {
		//go export.Consume(makeTcpStatsGaugeGroup(i))
		//time.Sleep(1 * time.Second)
	}

	wg.Wait()
}
