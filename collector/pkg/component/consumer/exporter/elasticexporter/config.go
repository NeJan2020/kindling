package elasticexporter

import (
	"strings"

	"github.com/Kindling-project/kindling/collector/pkg/model"
	"github.com/Kindling-project/kindling/collector/pkg/model/constlabels"
	es7 "github.com/elastic/go-elasticsearch/v7"
)

const FilterAll = "all"

type Config struct {
	EnableBulk bool        `mapstructure:"enable_bulk"`
	ESIndex    *ESIndex    `mapstructure:"es_index"`
	ESConfig   *es7.Config `mapstructure:"es_config"`
	// 1. namespace/workloadname: kindling/kindling-agent
	// 2. namespace: kindling
	SrcFilters []string `mapstructure:"src_filters"`
	DstFilters []string `mapstructure:"dst_filters"`
}

type ESIndex struct {
	TraceIndex  string `mapstructure:"trace_index"`
	MetricIndex string `mapstructure:"metric_index"`
}

type Filter interface {
	IsDrop(*model.DataGroup) bool
}

type Filters struct {
	Namespace    string
	WorkloadName string

	IsSrcFilter bool
}

func NewFilter(rule string, isSrcFilter bool) *Filters {
	ruleInfo := strings.Split(rule, "/")
	if len(ruleInfo) > 1 {
		// namespace / workload
		return &Filters{
			Namespace:    ruleInfo[0],
			WorkloadName: ruleInfo[1],
			IsSrcFilter:  isSrcFilter,
		}
	} else {
		return &Filters{
			Namespace:    rule,
			WorkloadName: FilterAll,
			IsSrcFilter:  isSrcFilter,
		}
	}
}

func (f *Filters) IsDrop(data *model.DataGroup) bool {
	var namespace, workload string
	if f.IsSrcFilter {
		namespace = data.Labels.GetStringValue(constlabels.SrcNamespace)
		workload = data.Labels.GetStringValue(constlabels.SrcWorkloadName)
	} else {
		namespace = data.Labels.GetStringValue(constlabels.DstNamespace)
		workload = data.Labels.GetStringValue(constlabels.DstWorkloadName)
	}
	if namespace != f.Namespace {
		return false
	}
	if workload != f.WorkloadName {
		return false
	}
	return true
}

func DefaultConfig() *Config {
	return &Config{
		ESIndex: &ESIndex{
			TraceIndex:  DefaultTraceIndex,
			MetricIndex: DefaultMetricIndex,
		},
		EnableBulk: false,
		// Default Config can not be used!
		ESConfig: &es7.Config{},
	}
}
