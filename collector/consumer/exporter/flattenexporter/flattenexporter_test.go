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
			{
				constvalues.ResponseIo,
				1234567891,
			},
			{
				constvalues.RequestTotalTime,
				int64(i),
			},
			{
				constvalues.RequestIo,
				4500,
			},
			{
				constvalues.RequestCount,
				4500,
			},
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

	gaugesGroup.Labels.AddStringValue(constlabels.Protocol, "http")
	gaugesGroup.Labels.AddStringValue(constlabels.StatusCode, "200")

	// Topology data preferentially use D Nat Ip and D Nat Port
	gaugesGroup.Labels.AddStringValue(constlabels.DstIp, "test-DnatIp")
	gaugesGroup.Labels.AddIntValue(constlabels.DstPort, 8081)
	return gaugesGroup
}

func makeAggNetGaugeGroup(i int) *model.GaugeGroup {
	gaugesGroup := &model.GaugeGroup{
		Name: constnames.AggregatedNetRequestGaugeGroup,
		Values: []*model.Gauge{
			{
				constvalues.ResponseIo,
				1234567891,
			},
			{
				constvalues.RequestTotalTime,
				int64(i),
			},
			{
				constvalues.RequestIo,
				4500,
			},
			{
				constvalues.RequestCount,
				4500,
			},
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

	gaugesGroup.Labels.AddStringValue(constlabels.Protocol, "http")
	gaugesGroup.Labels.AddStringValue(constlabels.StatusCode, "200")

	// Topology data preferentially use D Nat Ip and D Nat Port
	gaugesGroup.Labels.AddStringValue(constlabels.DstIp, "test-DnatIp")
	gaugesGroup.Labels.AddIntValue(constlabels.DstPort, 8081)
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
	for i := 0; i < 100; i++ {
		go export.Consume(makeSingleGaugeGroup(i))
		//time.Sleep(1 * time.Second)
	}
	//go export.Consume(makeSingleGaugeGroup(10))
	//go export.Consume(makeSingleGaugeGroup(11))
	//go export.Consume(makeSingleGaugeGroup(12))
	//go export.Consume(makeSingleGaugeGroup(13))
	//go export.Consume(makeAggNetGaugeGroup(10))
	//go export.Consume(makeAggNetGaugeGroup(11))
	//go export.Consume(makeAggNetGaugeGroup(12))
	//go export.Consume(makeAggNetGaugeGroup(13))
	/*	for i := 0; i < 10; i++ {
		go export.Consume(makeAggNetGaugeGroup(i))
		time.Sleep(1 * time.Second)
	}*/
	wg.Wait()
}
