package factory

type config struct {
	urlClusteringMethod string
	extractHost         bool
}

func newDefaultConfig() *config {
	return &config{
		urlClusteringMethod: "alphabet",
		extractHost:         false,
	}
}

type Option func(cfg *config)

func WithUrlClusteringMethod(urlClusteringMethod string) Option {
	return func(cfg *config) {
		cfg.urlClusteringMethod = urlClusteringMethod
	}
}

func WithExtractHost(enable bool) Option {
	return func(cfg *config) {
		cfg.extractHost = enable
	}
}
