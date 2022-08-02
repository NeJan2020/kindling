package kubernetes

import "time"

// config contains optional settings for connecting to kubernetes.
type config struct {
	KubeAuthType  AuthType
	KubeConfigDir string
	// GraceDeletePeriod controls the delay interval after receiving delete event.
	// The unit is seconds, and the default value is 60 seconds.
	// Should not be lower than 30 seconds.
	GraceDeletePeriod time.Duration

	// DSFRule
	DSFConfig *DSFConfig `mapstructure:"dsf_config"`
}

type DSFConfig struct {
	Enable           bool          `mapstructure:"enable"`
	ConfigServerAddr string        `mapstructure:"config_server_addr"`
	InitEndpoint     string        `mapstructure:"init_endpoint"`
	UpdateEndpoint   string        `mapstructure:"update_endpoint"`
	SyncInterval     time.Duration `mapstructure:"sync_interval"`
	EnableDebug      bool          `mapstructure:"debug"`
}

type Option func(cfg *config)

// WithAuthType sets way of authenticating kubernetes api-server
// Supported AuthTypeNone, AuthTypeServiceAccount, AuthTypeKubeConfig
func WithAuthType(authType AuthType) Option {
	return func(cfg *config) {
		cfg.KubeAuthType = authType
	}
}

// WithKubeConfigDir sets the directory where the file "kubeconfig" is stored
func WithKubeConfigDir(dir string) Option {
	return func(cfg *config) {
		cfg.KubeConfigDir = dir
	}
}

// WithGraceDeletePeriod sets the graceful period of deleting Pod's metadata
// after receiving "delete" event from client-go.
func WithGraceDeletePeriod(interval int) Option {
	return func(cfg *config) {
		cfg.GraceDeletePeriod = time.Duration(interval) * time.Second
	}
}

func WithDSFConfig(dsfCfg *DSFConfig) Option {
	return func(cfg *config) {
		cfg.DSFConfig = dsfCfg
	}
}
