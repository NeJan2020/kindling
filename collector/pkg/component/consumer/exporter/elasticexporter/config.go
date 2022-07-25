package elasticexporter

import es7 "github.com/elastic/go-elasticsearch/v7"

type Config struct {
	EnableBulk bool        `mapstructure:"enable_bulk"`
	ESIndex    *ESIndex    `mapstructure:"es_index"`
	ESConfig   *es7.Config `mapstructure:"es_config"`
}

type ESIndex struct {
	TraceIndex  string `mapstructure:"trace_index"`
	MetricIndex string `mapstructure:"metric_index"`
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
