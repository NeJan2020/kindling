package metadataclient

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/Kindling-project/kindling/collector/pkg/metadata/kubernetes"
	"github.com/Kindling-project/kindling/collector/pkg/metadata/metaprovider/api"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

type Client struct {
	// tls wrap
	cli      http.Client
	endpoint string
}

func NewMetaDataWrapperClient(endpoint string) *Client {
	return &Client{
		cli:      *createHTTPClient(),
		endpoint: endpoint,
	}
}

func (c *Client) ListAndWatch() error { //req api.MetaDataRequest,
	// handler cache.ResourceEventHandler,
	resp, err := c.cli.Get(c.endpoint)
	if err != nil {
		return err
	}
	reader := bufio.NewReaderSize(resp.Body, 1024*32)
	b, _ := reader.ReadBytes('\n')
	// var listData api.MetaData
	// json.Unmarshal(b, &listData)
	// cache := kubernetes.K8sMetaDataCache{}
	listVO := api.ListVO{}
	err = json.Unmarshal(b, &listVO)
	kubernetes.SetPreprocessingMetaDataCache(listVO.Cache, listVO.GlobalNodeInfo, listVO.GlobalServiceInfo, listVO.GlobalRsInfo)
	if err != nil {
		log.Printf("list failed,err %v", err)
	}
	for {
		b, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				log.Printf("remote server send unexpected EOF,shutting down,err:%v", err)
				break
			} else {
				log.Printf("receive unexpected error durning watch ,err:%v", err)
				break
			}
		}
		var resp api.MetaDataVO
		err = json.Unmarshal(b, &resp)
		if err != nil {
			log.Printf("parse response failed ,err:%v", err)
			continue
		}
		convertRespToEventAndApply(resp)
	}
	return nil
}

func convertRespToEventAndApply(resp api.MetaDataVO) {
	switch resp.Operation {
	case "add":
		if resp.NewObj == nil {
			log.Printf("convert k8sMetadata falled, err: newObj is nill in add Event")
			return
		}

		switch resp.Type {
		case "node":
			obj := corev1.Node{}
			err := json.Unmarshal(resp.NewObj, &obj)
			if err != nil {
				log.Printf("convert k8sMetadata falled, err: %v", err)
				return
			}
			log.Printf("Debug: %s: %s ,Name: %s", resp.Operation, resp.Type, obj.Name)
			kubernetes.AddNode(&obj)
		case "pod":
			obj := corev1.Pod{}
			err := json.Unmarshal(resp.NewObj, &obj)
			if err != nil {
				log.Printf("convert k8sMetadata falled, err: %v", err)
				return
			}
			log.Printf("Debug: %s: %s ,Name: %s", resp.Operation, resp.Type, obj.Name)
			kubernetes.PodAdd(&obj)
		case "service":
			obj := corev1.Service{}
			err := json.Unmarshal(resp.NewObj, &obj)
			if err != nil {
				log.Printf("convert k8sMetadata falled, err: %v", err)
				return
			}
			log.Printf("Debug: %s: %s ,Name: %s", resp.Operation, resp.Type, obj.Name)
			kubernetes.AddService(&obj)
		case "rs":
			obj := appv1.ReplicaSet{}
			err := json.Unmarshal(resp.NewObj, &obj)
			if err != nil {
				log.Printf("convert k8sMetadata falled, err: %v", err)
				return
			}
			log.Printf("Debug: %s: %s ,Name: %s", resp.Operation, resp.Type, obj.Name)
			kubernetes.AddReplicaSet(&obj)
		}
	case "update":
		if resp.NewObj == nil || resp.OldObj == nil {
			log.Printf("convert k8sMetadata falled, err: newObj or oldObj is nill in update Event")
			return
		}

		switch resp.Type {
		case "node":
			newObj := corev1.Node{}
			err := json.Unmarshal(resp.NewObj, &newObj)
			if err != nil {
				log.Printf("convert k8sMetadata falled, err: %v", err)
				return
			}
			oldObj := corev1.Node{}
			err = json.Unmarshal(resp.OldObj, &oldObj)
			if err != nil {
				log.Printf("convert k8sMetadata falled, err: %v", err)
				return
			}
			kubernetes.UpdateNode(&oldObj, &newObj)
		case "pod":
			newObj := corev1.Pod{}
			err := json.Unmarshal(resp.NewObj, &newObj)
			if err != nil {
				log.Printf("convert k8sMetadata falled, err: %v", err)
				return
			}
			oldObj := corev1.Pod{}
			err = json.Unmarshal(resp.OldObj, &oldObj)
			if err != nil {
				log.Printf("convert k8sMetadata falled, err: %v", err)
				return
			}
			log.Printf("Debug: %s: %s ,Name: %s", resp.Operation, resp.Type, newObj.Name)
			kubernetes.PodUpdate(&oldObj, &newObj)
		case "service":
			newObj := corev1.Service{}
			err := json.Unmarshal(resp.NewObj, &newObj)
			if err != nil {
				log.Printf("convert k8sMetadata falled, err: %v", err)
				return
			}
			oldObj := corev1.Service{}
			err = json.Unmarshal(resp.OldObj, &oldObj)
			if err != nil {
				log.Printf("convert k8sMetadata falled, err: %v", err)
				return
			}
			kubernetes.UpdateService(&oldObj, &newObj)
		case "rs":
			newObj := appv1.ReplicaSet{}
			err := json.Unmarshal(resp.NewObj, &newObj)
			if err != nil {
				log.Printf("convert k8sMetadata falled, err: %v", err)
				return
			}
			oldObj := appv1.ReplicaSet{}
			err = json.Unmarshal(resp.OldObj, &oldObj)
			if err != nil {
				log.Printf("convert k8sMetadata falled, err: %v", err)
				return
			}
			kubernetes.UpdateReplicaSet(&oldObj, &newObj)
		}
	case "delete":
		if resp.OldObj == nil {
			log.Printf("convert k8sMetadata falled, err: newObj is nill in add Event")
			return
		}

		switch resp.Type {
		case "node":
			obj := corev1.Node{}
			err := json.Unmarshal(resp.OldObj, &obj)
			if err != nil {
				log.Printf("convert k8sMetadata falled, err: %v", err)
				return
			}
			log.Printf("Debug: %s: %s ,Name: %s", resp.Operation, resp.Type, obj.Name)
			kubernetes.DeleteNode(&obj)
		case "pod":
			obj := corev1.Pod{}
			err := json.Unmarshal(resp.OldObj, &obj)
			if err != nil {
				log.Printf("convert k8sMetadata falled, err: %v", err)
				return
			}
			log.Printf("Debug: %s: %s ,Name: %s", resp.Operation, resp.Type, obj.Name)
			kubernetes.PodDelete(&obj)
		case "service":
			obj := corev1.Service{}
			err := json.Unmarshal(resp.OldObj, &obj)
			if err != nil {
				log.Printf("convert k8sMetadata falled, err: %v", err)
				return
			}
			log.Printf("Debug: %s: %s ,Name: %s", resp.Operation, resp.Type, obj.Name)
			kubernetes.DeleteService(&obj)
		case "rs":
			obj := appv1.ReplicaSet{}
			err := json.Unmarshal(resp.OldObj, &obj)
			if err != nil {
				log.Printf("convert k8sMetadata falled, err: %v", err)
				return
			}
			log.Printf("Debug: %s: %s ,Name: %s", resp.Operation, resp.Type, obj.Name)
			kubernetes.DeleteReplicaSet(&obj)
		}
	}
}

func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
		},
	}
	return client
}
