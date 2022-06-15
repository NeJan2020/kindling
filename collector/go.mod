module github.com/Kindling-project/kindling/collector

go 1.16

require (
	github.com/DataDog/ebpf v0.0.0-20220301203322-3fc9ab3b8daf
	github.com/cenkalti/backoff/v4 v4.1.1
	github.com/elastic/gosigar v0.14.2
	github.com/florianl/go-conntrack v0.3.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hashicorp/golang-lru v0.5.4
	github.com/jaegertracing/jaeger v1.33.0
	github.com/mdlayher/netlink v1.6.0
	github.com/orcaman/concurrent-map v0.0.0-20210501183033-44dafcb38ecc
	github.com/pebbe/zmq4 v1.2.7
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.12.1
	github.com/rs/cors v1.8.2
	github.com/shirou/gopsutil v3.21.11+incompatible
	github.com/spf13/cast v1.4.1
	github.com/spf13/viper v1.10.1
	github.com/stretchr/testify v1.7.1
	github.com/tklauser/go-sysconf v0.3.10 // indirect
	github.com/vishvananda/netns v0.0.0-20211101163701-50045581ed74
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	go.opencensus.io v0.23.0
	go.opentelemetry.io/otel v1.2.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.25.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.2.0
	go.opentelemetry.io/otel/exporters/prometheus v0.25.0
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v0.25.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.2.0
	go.opentelemetry.io/otel/metric v0.25.0
	go.opentelemetry.io/otel/sdk v1.2.0
	go.opentelemetry.io/otel/sdk/export/metric v0.25.0
	go.opentelemetry.io/otel/sdk/metric v0.25.0
	go.opentelemetry.io/otel/trace v1.2.0
	go.uber.org/multierr v1.7.0
	go.uber.org/zap v1.21.0
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f
	golang.org/x/sys v0.0.0-20220128215802-99c3d69c2c27
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	k8s.io/api v0.21.5
	k8s.io/apimachinery v0.21.5
	k8s.io/client-go v0.21.5
)
