package api

import (
	"encoding/json"

	"github.com/Kindling-project/kindling/collector/pkg/metadata/kubernetes"
)

type MetaDataService interface {
	ListAndWatch()
}

type MetaDataRsyncResponse struct {
	Type      string // pod,rs,node,service
	Operation string // add,update,delete
	NewObj    interface{}
	OldObj    interface{}
}

type MetaDataVO struct {
	Type      string // pod,rs,node,service
	Operation string // add,update,delete
	NewObj    json.RawMessage
	OldObj    json.RawMessage
}

type ListVO struct {
	Cache             *kubernetes.K8sMetaDataCache
	GlobalNodeInfo    *kubernetes.NodeMap
	GlobalRsInfo      *kubernetes.ReplicaSetMap
	GlobalServiceInfo *kubernetes.ServiceMap
}
