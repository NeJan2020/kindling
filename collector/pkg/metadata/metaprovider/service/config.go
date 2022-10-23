package service

import "github.com/Kindling-project/kindling/collector/pkg/metadata/kubernetes"

type Config struct {
	KubeAuthType  kubernetes.AuthType
	KubeConfigDir string
}
