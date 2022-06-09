package flattenexporter

import (
	"strconv"
	"sync"
	"testing"

	"github.com/Kindling-project/kindling/collector/component"
	"github.com/Kindling-project/kindling/collector/model"
	"github.com/Kindling-project/kindling/collector/model/constlabels"
	"github.com/Kindling-project/kindling/collector/model/constnames"
	"github.com/Kindling-project/kindling/collector/model/constvalues"
	"github.com/spf13/viper"
)

func makeSingleDataGroup(i int) *model.DataGroup {
	MetricsGroup := &model.DataGroup{
		Name: constnames.SingleNetRequestMetricGroup,
		Metrics: []*model.Metric{
			model.NewIntMetric(constvalues.ResponseIo, 123456789),
			model.NewIntMetric(constvalues.RequestIo, 1234567891),
			model.NewIntMetric(constvalues.RequestTotalTime, int64(i+9)),
			model.NewIntMetric(constvalues.ConnectTime, 4500),
			model.NewIntMetric(constvalues.RequestSentTime, 4500),
			model.NewIntMetric(constvalues.WaitingTtfbTime, 4500),
			model.NewIntMetric(constvalues.ContentDownloadTime, 4500),
			model.NewIntMetric(constvalues.RequestCount, 4500),
			//{Name: connect_time, Value: 0}
			//{Name: request_sent_time, Value: 9517}
			//{Name: waiting_ttfb_time, Value: 499799900}
			//{Name: content_download_time, Value: 79743}
			//{Name: request_total_time, Value: 499889160}
			//{Name: request_io, Value: 71}
			//{Name: response_io, Value: 22}
		},
		Labels:    model.NewAttributeMap(),
		Timestamp: 19900909090111111,
	}
	MetricsGroup.Labels.AddBoolValue(constlabels.IsSlow, true)
	MetricsGroup.Labels.AddBoolValue(constlabels.IsError, true)
	MetricsGroup.Labels.AddBoolValue(constlabels.IsServer, true)
	MetricsGroup.Labels.AddStringValue(constlabels.SrcIp, "test-SrcIp"+strconv.Itoa(i))
	MetricsGroup.Labels.AddIntValue(constlabels.SrcPort, 8080)
	MetricsGroup.Labels.AddStringValue(constlabels.SrcNode, "test-SrcNode"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.SrcNamespace, "test-SrcNamespace"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.SrcWorkloadKind, "test-SrcWorkloadKind"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.SrcWorkloadName, "test-SrcWorkloadName"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.SrcService, "test-SrcService"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.SrcPod, "test-SrcPod"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.DstIp, "test-dstIp"+strconv.Itoa(i))
	MetricsGroup.Labels.AddIntValue(constlabels.DstPort, 8090)
	MetricsGroup.Labels.AddStringValue(constlabels.DnatIp, "test-dnatIp"+strconv.Itoa(i))
	MetricsGroup.Labels.AddIntValue(constlabels.DnatPort, 8091)
	MetricsGroup.Labels.AddStringValue(constlabels.DstNode, "test-DstNode"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.DstNamespace, "test-DstNamespace"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.DstWorkloadName, "test-DstWorkloadName"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.DstWorkloadKind, "test-DstWorkloadKind"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.DstService, "test-DstService"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.DstPod, "test-DstPod"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.HttpMethod, "test-HttpMethod"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.HttpUrl, "test-HttpUrl"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.HttpRequestPayload, "GET / HTTP/1.1\\r\\nHost: :8080\\r\\nUser-Agent: Go-http-client/1.1\\r\\nAccept-Encoding: gz")
	MetricsGroup.Labels.AddStringValue(constlabels.HttpRequestPayload, "HTTP/1.1 404 Not Found\\r\\nContent-Type: text/plain; charset=utf-8\\r\\nX-Content-Type-")
	MetricsGroup.Labels.AddStringValue(constlabels.SrcContainerId, "test-SrcContainerId"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.SrcContainer, "test-SrcContainerName"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.DstContainerId, "test-DstContainerId"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.DstContainer, "test-DstContainerName"+strconv.Itoa(i))
	MetricsGroup.Labels.AddIntValue(constlabels.HttpStatusCode, 299)
	MetricsGroup.Labels.AddIntValue(constlabels.KafkaErrorCode, 298)
	MetricsGroup.Labels.AddIntValue(constlabels.SqlErrCode, 297)
	MetricsGroup.Labels.AddIntValue(constlabels.Pid, int64(i))
	MetricsGroup.Labels.AddStringValue(constlabels.Protocol, "dns")
	MetricsGroup.Labels.AddStringValue(constlabels.ContentKey, "/url")
	MetricsGroup.Labels.AddIntValue(constlabels.DnsRcode, int64(i+200))
	MetricsGroup.Labels.AddStringValue(constlabels.DnsDomain, "DNS_DOMAIN")
	return MetricsGroup
}

func makeAggNetDataGroup(i int) *model.DataGroup {
	MetricsGroup := &model.DataGroup{
		Name: constnames.AggregatedNetRequestMetricGroup,
		Metrics: []*model.Metric{
			model.NewHistogramMetric(constvalues.RequestTotalTime, &model.Histogram{
				Sum:                10000*4999 + 15000,
				Count:              10000,
				ExplicitBoundaries: []int64{0, 100, 200, 500, 1000, 2000, 5000, 10000},
				BucketCounts:       []uint64{0, 100, 200, 500, 1000, 2000, 5000, 10000},
			}),
			model.NewIntMetric(constvalues.ResponseIo, 1234567891),
			model.NewIntMetric(constvalues.RequestIo, 4500),
		},
		Labels:    model.NewAttributeMap(),
		Timestamp: 1653630590614288940,
	}
	MetricsGroup.Labels.AddStringValue(constlabels.SrcNode, "test-SrcNode"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.SrcNamespace, "test-SrcNamespace"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.SrcPod, "test-SrcPod"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.SrcWorkloadName, "test-SrcWorkloadName"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.SrcWorkloadKind, "test-SrcWorkloadKind"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.SrcService, "test-SrcService"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.SrcIp, "test-SrcIp"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.DstNode, "test-DstNode"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.DstNamespace, "test-DstNamespace"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.DstPod, "test-DstPod"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.DstWorkloadName, "test-DstWorkloadName"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.DstWorkloadKind, "test-DstWorkloadKind"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.DstService, "test-DstService"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.SrcContainer, "test-SrcContainer"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.SrcContainerId, "test-SrcContainerId"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.DstContainer, "test-SrcContainer"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.DstContainerId, "test-SrcContainerId"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.Protocol, "kafka")
	//MetricsGroup.Labels.AddStringValue(constlabels.Protocol, "http")
	MetricsGroup.Labels.AddIntValue(constlabels.HttpStatusCode, 498)
	MetricsGroup.Labels.AddBoolValue(constlabels.IsSlow, true)
	MetricsGroup.Labels.AddBoolValue(constlabels.IsError, true)
	MetricsGroup.Labels.AddBoolValue(constlabels.IsServer, true)
	MetricsGroup.Labels.AddStringValue(constlabels.ContentKey, "/http/url")
	// Topology data preferentially use D Nat Ip and D Nat Port
	MetricsGroup.Labels.AddStringValue(constlabels.DstIp, "test-DstIp")
	MetricsGroup.Labels.AddIntValue(constlabels.DstPort, 8081)
	MetricsGroup.Labels.AddStringValue(constlabels.KafkaTopic, "npm_detail_topology_request")
	MetricsGroup.Labels.AddStringValue(constlabels.DnatIp, "test-DnatIp"+strconv.Itoa(i))
	MetricsGroup.Labels.AddIntValue(constlabels.DnatPort, 8082)
	MetricsGroup.Labels.AddIntValue(constlabels.Pid, 111)
	return MetricsGroup
}
func makeTcpStatsDataGroup(i int) *model.DataGroup {
	MetricsGroup := &model.DataGroup{
		Name: constnames.TcpStatsMetricGroup,
		Metrics: []*model.Metric{
			model.NewIntMetric("Established", int64(i)),
			model.NewIntMetric("SynSent", int64(i)),
			model.NewIntMetric("SynRecv", int64(i)),
			model.NewIntMetric("FinWait1", int64(i)),
			model.NewIntMetric("FinWait2", int64(i)),
			model.NewIntMetric("TimeWait", int64(i)),
			model.NewIntMetric("Close", int64(i)),
			model.NewIntMetric("CloseWait", int64(i)),
			model.NewIntMetric("LastAck", int64(i)),
			model.NewIntMetric("Listen", int64(i)),
			model.NewIntMetric("Closing", int64(i)),
		},
		Labels:    model.NewAttributeMap(),
		Timestamp: 19900909090,
	}
	MetricsGroup.Labels.AddStringValue("mode", "tcp4")
	MetricsGroup.Labels.AddStringValue("container", "test-container"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue("container_id", "test-container_id"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue("namespace", "test-namespace"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue("node", "test-node"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue("node_ip", "test-node_ip"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue("pod", "test-pod"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue("service", "test-elasticsearch-svc"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue("workload_kind", "test-statefulset"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue("workload_name", "test-workload_name"+strconv.Itoa(i))
	return MetricsGroup
}

func makePageFaultDataGroup(i int) *model.DataGroup {
	MetricsGroup := &model.DataGroup{
		Name: constnames.PgftGaugeGroupName,
		Metrics: []*model.Metric{
			model.NewIntMetric("kindling_pagefault_major_total", int64(i)),
			model.NewIntMetric("kindling_pagefault_minor_total", int64(i)),
		},
		Labels:    model.NewAttributeMap(),
		Timestamp: 19900909090,
	}
	MetricsGroup.Labels.AddStringValue("tid", "test-tid"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue("pid", "test-pid"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue("container_id", "test-container_id"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue("namespace", "test-namespace"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue("node", "test-node"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue("node_ip", "test-node_ip"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue("pod", "test-pod"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue("service", "test-elasticsearch-svc"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue("workload_kind", "test-statefulset"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue("workload_name", "test-workload_name"+strconv.Itoa(i))
	return MetricsGroup
}

func makeTcpMetricDataGroup(i int) *model.DataGroup {
	MetricsGroup := &model.DataGroup{
		Name: constnames.TcpConnectMetricGroupName,
		Metrics: []*model.Metric{
			model.NewIntMetric(constlabels.KindlingTcpConnectTotal, int64(i)),
			model.NewHistogramMetric(constlabels.KindlingTcpConnectDurationNanoseconds, &model.Histogram{
				Sum:                15000,
				Count:              10000,
				ExplicitBoundaries: []int64{0, 100, 200, 500, 1000, 2000, 5000, 10000},
				BucketCounts:       []uint64{0, 100, 200, 500, 1000, 2000, 5000, 10000},
			}),
		},
		Labels:    model.NewAttributeMap(),
		Timestamp: 19900909090,
	}

	MetricsGroup.Labels.AddStringValue(constlabels.DstContainer, "test-DstContainer"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.DstContainerId, "test-DstContainerId"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.DstIp, "test-DnatIp")
	MetricsGroup.Labels.AddIntValue(constlabels.DstPort, 8081)
	MetricsGroup.Labels.AddStringValue(constlabels.DstNode, "test-DstNode"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.DstNodeIp, "test-DstNodeIp"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.DstNamespace, "test-DstNamespace"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.DstPod, "test-DstPod"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.DstWorkloadName, "test-DstWorkloadName"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.DstWorkloadKind, "test-DstWorkloadKind"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.DstService, "test-DstService"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.SrcContainer, "test-SrcContainer"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.SrcContainerId, "test-SrcContainerId"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.SrcIp, "test-SrcIp"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.SrcPort, "test-SrcPort"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.SrcNode, "test-SrcNode"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.SrcNodeIp, "test-SrcNodeIp"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.SrcNamespace, "test-SrcNamespace"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.SrcPod, "test-SrcPod"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.SrcWorkloadName, "test-SrcWorkloadName"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.SrcWorkloadKind, "test-SrcWorkloadKind"+strconv.Itoa(i))
	MetricsGroup.Labels.AddStringValue(constlabels.SrcService, "test-SrcService"+strconv.Itoa(i))
	MetricsGroup.Labels.AddBoolValue(constlabels.IsServer, false)
	MetricsGroup.Labels.AddBoolValue(constlabels.Success, false)
	return MetricsGroup
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
		//go export.Consume(makeSingleDataGroup(i))
		//time.Sleep(1 * time.Second)
		//go export.Consume(makePageFaultDataGroup(i))
	}

	for i := 1; i < 2; i++ {
		//go export.Consume(makeSingleDataGroup(i))
		//go export.Consume(makeAggNetDataGroup(i))
		//go export.Consume(makeTcpStatsDataGroup(i))
		go export.Consume(makePageFaultDataGroup(i))

		//go export.Consume(makeTcpMetricDataGroup(i))
	}
	for i := 0; i < 10; i++ {
		//go export.Consume(makeTcpStatsDataGroup(i))
		//time.Sleep(1 * time.Second)
	}

	wg.Wait()
}
