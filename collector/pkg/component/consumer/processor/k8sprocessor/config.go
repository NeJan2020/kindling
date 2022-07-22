package k8sprocessor

import (
	"time"

	"github.com/Kindling-project/kindling/collector/pkg/metadata/kubernetes"
)

type Config struct {
	KubeAuthType  kubernetes.AuthType `mapstructure:"kube_auth_type"`
	KubeConfigDir string              `mapstructure:"kube_config_dir"`
	// GraceDeletePeriod controls the delay interval after receiving delete event.
	// The unit is seconds, and the default value is 60 seconds.
	// Should not be lower than 30 seconds.
	GraceDeletePeriod int `mapstructure:"grace_delete_period"`
	// Set "Enable" false if you want to run the agent in the non-Kubernetes environment.
	// Otherwise, the agent will panic if it can't connect to the API-server.
	Enable bool `mapstructure:"enable"`

	// DSFRule
	DSFConfig *kubernetes.DSFConfig `mapstructure:"dsf_config"`
}

var DefaultConfig Config = Config{
	KubeAuthType:      "serviceAccount",
	KubeConfigDir:     "~/.kube/config",
	GraceDeletePeriod: 60,
	Enable:            true,
	DSFConfig: &kubernetes.DSFConfig{
		Enable:         false,
		InitEndpoint:   "/hcmine/config/dsfInit",
		UpdateEndpoint: "/hcmine/config/dsfUpdate",
		SyncInterval:   5 * time.Second,
		EnableDebug:    false,
	},
}
