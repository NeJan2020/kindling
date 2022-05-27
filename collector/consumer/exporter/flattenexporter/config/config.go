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

// Package config defines the data models for entities. This file defines the
// models for configuration format. The defined entities are:
// Config (the top-level structure), Receivers, Exporters, Processors, Pipelines.
//
// Receivers, Exporters and Processors typically have common configuration settings, however
// sometimes specific implementations will have extra configuration settings.
// This requires the configuration data for these entities to be polymorphic.
//
// To satisfy these requirements we declare interfaces Receiver, Exporter, Processor,
// which define the behavior. We also provide helper structs ReceiverSettings, ExporterSettings,
// ProcessorSettings, which define the common settings and un-marshaling from config files.
//
// Specific Receivers/Exporters/Processors are expected to at the minimum implement the
// corresponding interface and if they have additional settings they must also extend
// the corresponding common settings struct (the easiest approach is to embed the common struct).
package config

import (
	"errors"
	"fmt"
)

var (
	errMissingExporters        = errors.New("no enabled exporters specified in config")
	errMissingReceivers        = errors.New("no enabled receivers specified in config")
	errMissingServicePipelines = errors.New("service must have at least one pipeline")
)

// Config defines the configuration for the various elements of collector or agent.
type Config struct {
	Exporters
	Extensions
	Service
}

var _ validatable = (*Config)(nil)

// Validate returns an error if the config is invalid.
//
// This function performs basic validation of configuration. There may be more subtle
// invalid cases that we currently don't check for but which we may want to add in
// the future (e.g. disallowing receiving and exporting on the same endpoint).
func (cfg *Config) Validate() error {
	// Currently there is no default exporter enabled.
	// The configuration must specify at least one exporter to be valid.
	if len(cfg.Exporters) == 0 {
		return errMissingExporters
	}

	// Validate the exporter configuration.
	for exp, expCfg := range cfg.Exporters {
		if err := expCfg.Validate(); err != nil {
			return fmt.Errorf("exporter \"%s\" has invalid configuration: %w", exp, err)
		}
	}

	// Validate the extension configuration.
	for ext, extCfg := range cfg.Extensions {
		if err := extCfg.Validate(); err != nil {
			return fmt.Errorf("extension \"%s\" has invalid configuration: %w", ext, err)
		}
	}

	// Check that all enabled extensions in the service are configured
	if err := cfg.validateServiceExtensions(); err != nil {
		return err
	}

	// Check that all pipelines have at least one receiver and one exporter, and they reference
	// only configured components.
	return cfg.validateServicePipelines()
}

func (cfg *Config) validateServiceExtensions() error {
	// Validate extensions.
	for _, ref := range cfg.Service.Extensions {
		// Check that the name referenced in the Service extensions exists in the top-level extensions
		if cfg.Extensions[ref] == nil {
			return fmt.Errorf("service references extension %q which does not exist", ref)
		}
	}

	return nil
}

func (cfg *Config) validateServicePipelines() error {
	// Must have at least one pipeline.
	if len(cfg.Service.Pipelines) == 0 {
		return errMissingServicePipelines
	}

	// Validate pipelines.
	for _, pipeline := range cfg.Service.Pipelines {
		for _, mtl := range pipeline.MTLPipelines {
			// Validate pipeline processor name references

			// Validate pipeline has at least one exporter
			if len(mtl.Exporters) == 0 {
				return fmt.Errorf("pipeline %q must have at least one exporter", mtl.Name)
			}

			// Validate pipeline exporter name references.
			for _, ref := range mtl.Exporters {
				// Check that the name referenced in the pipeline's Exporters exists in the top-level Exporters
				if cfg.Exporters[ref] == nil {
					return fmt.Errorf("pipeline %q references exporter %q which does not exist", mtl.Name, ref)
				}
			}
		}
	}
	return nil
}

// Service defines the configurable components of the service.
type Service struct {
	// Extensions is the ordered list of extensions configured for the service.
	Extensions []ComponentID

	// Pipelines is the set of data pipelines configured for the service.
	Pipelines Pipelines
}

// Type is the component type as it is used in the config.
type Type string

// validatable defines the interface for the configuration validation.
type validatable interface {
	// Validate validates the configuration and returns an error if invalid.
	Validate() error
}

// CustomUnmarshable defines an optional interface for custom configuration unmarshaling.
// A configuration struct can implement this interface to override the default unmarshaling.
type CustomUnmarshable interface {
	// Unmarshal is a function that un-marshals a Parser into the unmarshable struct in a custom way.
	// componentSection *Parser
	//   The config for this specific component. May be nil or empty if no config available.
	Unmarshal(componentSection *Parser) error
}

// DataType is the data type that is supported for collection. We currently support
// collecting metrics, traces and logs, this can expand in the future.
type DataType string

// Currently supported data types. Add new data types here when new types are supported in the future.
const (
	// TracesDataType is the data type tag for traces.
	TracesDataType DataType = "traces"

	// MetricsDataType is the data type tag for metrics.
	MetricsDataType DataType = "metrics"

	// LogsDataType is the data type tag for logs.
	LogsDataType DataType = "logs"

	HcmineEventDataType DataType = "HcmineEvent"
)

// Pipeline defines a single pipeline.
type Pipeline struct {
	Name         string
	Receivers    []ComponentID
	Analyzers    []ComponentID
	MTLPipelines map[string]*MTLPipeline
}

type MTLPipeline struct {
	Name       string
	InputType  DataType
	Processors []ComponentID
	Exporters  []ComponentID
}

// Pipelines is a map of names to Pipelines.
type Pipelines map[string]*Pipeline

// MTLPipelines is a map of names to MTLPipelines.
type MTLPipelines map[string]*MTLPipeline