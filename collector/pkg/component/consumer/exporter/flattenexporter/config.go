// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package flattenexporter

import (
	"github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/config"
	"github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/config/confighttp"
	otlpcommon "github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/data/protogen/common/v1"
	"github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/internal/exporterhelper"
	"os"
	"time"
)

// Config defines configuration for flatten exporter.
type Config struct {
	MasterIp                      string                   `mapstructure:"master_ip"`
	BatchMaxSize                  int                      `mapstructure:"batch_max_size"`
	config.ExporterSettings       `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct
	confighttp.HTTPClientSettings `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct.
	exporterhelper.QueueSettings  `mapstructure:"sending_queue"`
	exporterhelper.RetrySettings  `mapstructure:"retry_on_failure"`

	// The URL to send traces to. If omitted the Endpoint + "/v1/traces" will be used.
	TracesEndpoint string `mapstructure:"traces_endpoint"`

	// The URL to send metrics to. If omitted the Endpoint + "/v1/metrics" will be used.
	MetricsEndpoint string `mapstructure:"metrics_endpoint"`

	// The URL to send logs to. If omitted the Endpoint + "/v1/logs" will be used.
	LogsEndpoint string `mapstructure:"logs_endpoint"`

	// The compression key for supported compression types within
	// collector. Currently the only supported mode is `gzip`.
	Compression string `mapstructure:"compression"`

	// Timeout sets the time after which a batch will be sent regardless of size.
	BatchTimeout time.Duration `mapstructure:"batch_timeout"`
	// SendBatchSize is the size of a batch which after hit, will trigger it to be sent.
	SendBatchSize uint32 `mapstructure:"send_batch_size"`

	SendBatchMaxSize uint32 `mapstructure:"send_batch_max_size"`
}

var _ config.Exporter = (*Config)(nil)

var serviceInstance *otlpcommon.Service

func (cfg *Config) GetServiceInstance() *otlpcommon.Service {
	if serviceInstance == nil {
		if hostName, err := os.Hostname(); err == nil {
			serviceInstance = &otlpcommon.Service{Job: "kindling", Instance: hostName, MasterIp: cfg.MasterIp}
		} else {
			serviceInstance = &otlpcommon.Service{Job: "kindling", Instance: "UnKnowHost", MasterIp: cfg.MasterIp}
		}
	}
	return serviceInstance
}
