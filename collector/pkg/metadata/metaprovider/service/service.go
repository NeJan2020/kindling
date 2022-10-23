package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/Kindling-project/kindling/collector/pkg/metadata/kubernetes"
	"github.com/Kindling-project/kindling/collector/pkg/metadata/metaprovider/api"
	"github.com/Kindling-project/kindling/collector/pkg/metadata/metaprovider/ioutil"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
)

type MetaDataWrapper struct {
	// TODO Scricube
	flushers []*Watcher

	// Signal
	stopCh chan struct{}
}

func NewMetaDataWrapper(config *Config) (*MetaDataWrapper, error) {
	mp := &MetaDataWrapper{
		flushers: make([]*Watcher, 0),
		stopCh:   make(chan struct{}),
	}

	// DEBUG Watcher Size
	// go func() {
	// 	ticker := time.NewTicker(5 * time.Second)
	// 	for {
	// 		select {
	// 		case <-ticker.C:
	// 			log.Printf("Debug: Watcher List Length: %d", len(mp.flushers))
	// 		}
	// 	}
	// }()

	var options []kubernetes.Option
	options = append(options, kubernetes.WithAuthType(config.KubeAuthType))
	options = append(options, kubernetes.WithKubeConfigDir(config.KubeConfigDir))
	options = append(options, kubernetes.WithGraceDeletePeriod(0))
	options = append(options, kubernetes.WithPodEventHander(cache.ResourceEventHandlerFuncs{
		AddFunc:    mp.AddPod,
		UpdateFunc: mp.UpdatePod,
		DeleteFunc: mp.DeletePod,
	}))
	options = append(options, kubernetes.WithServiceEventHander(cache.ResourceEventHandlerFuncs{
		AddFunc:    mp.AddSerivce,
		UpdateFunc: mp.UpdateSerivce,
		DeleteFunc: mp.DeleteSerivce,
	}))
	options = append(options, kubernetes.WithNodeEventHander(cache.ResourceEventHandlerFuncs{
		AddFunc:    mp.AddNode,
		UpdateFunc: mp.UpdateNode,
		DeleteFunc: mp.DeleteNode,
	}))
	options = append(options, kubernetes.WithServiceEventHander(cache.ResourceEventHandlerFuncs{
		AddFunc:    mp.AddSerivce,
		UpdateFunc: mp.UpdateSerivce,
		DeleteFunc: mp.DeleteSerivce,
	}))
	return mp, kubernetes.InitK8sHandler(options...)
}

//TODO Add Item
func (s *MetaDataWrapper) AddPod(obj interface{}) {
	kubernetes.PodAdd(obj)
	clearPodInfo(obj)
	s.boardcast(api.MetaDataRsyncResponse{
		Type:      "pod",
		Operation: "add",
		NewObj:    obj,
	})
}

func (s *MetaDataWrapper) AddRs(obj interface{}) {
	kubernetes.AddReplicaSet(obj)
	s.boardcast(api.MetaDataRsyncResponse{
		Type:      "rs",
		Operation: "add",
		NewObj:    obj,
	})
}

func (s *MetaDataWrapper) AddNode(obj interface{}) {
	kubernetes.AddNode(obj)
	s.boardcast(api.MetaDataRsyncResponse{
		Type:      "node",
		Operation: "add",
		NewObj:    obj,
	})
}

func (s *MetaDataWrapper) AddSerivce(obj interface{}) {
	kubernetes.AddService(obj)
	s.boardcast(api.MetaDataRsyncResponse{
		Type:      "service",
		Operation: "add",
		NewObj:    obj,
	})
}

func (s *MetaDataWrapper) UpdatePod(objOld interface{}, objNew interface{}) {
	kubernetes.PodUpdate(objOld, objNew)

	clearPodInfo(objOld)
	clearPodInfo(objNew)

	s.boardcast(api.MetaDataRsyncResponse{
		Type:      "pod",
		Operation: "update",
		OldObj:    objOld,
		NewObj:    objNew,
	})
}

func (s *MetaDataWrapper) UpdateNode(objOld interface{}, objNew interface{}) {
	kubernetes.UpdateNode(objOld, objNew)
	s.boardcast(api.MetaDataRsyncResponse{
		Type:      "node",
		Operation: "update",
		OldObj:    objOld,
		NewObj:    objNew,
	})
}

func (s *MetaDataWrapper) UpdateSerivce(objOld interface{}, objNew interface{}) {
	kubernetes.UpdateService(objOld, objNew)
	s.boardcast(api.MetaDataRsyncResponse{
		Type:      "service",
		Operation: "update",
		OldObj:    objOld,
		NewObj:    objNew,
	})
}

func clearPodInfo(objOld interface{}) {
	pod := objOld.(*corev1.Pod)
	//  Clear unnecesssay Message
	pod.ManagedFields = nil
	pod.Spec.Volumes = nil
	pod.Status.Conditions = nil
}

func (s *MetaDataWrapper) DeletePod(objOld interface{}) {
	kubernetes.PodDelete(objOld)
	clearPodInfo(objOld)

	s.boardcast(api.MetaDataRsyncResponse{
		Type:      "pod",
		Operation: "delete",
		OldObj:    objOld,
	})
}

func (s *MetaDataWrapper) DeleteRs(obj interface{}) {
	kubernetes.DeleteReplicaSet(obj)
	s.boardcast(api.MetaDataRsyncResponse{
		Type:      "rs",
		Operation: "delete",
		OldObj:    obj,
	})
}

func (s *MetaDataWrapper) DeleteNode(obj interface{}) {
	kubernetes.DeleteNode(obj)
	s.boardcast(api.MetaDataRsyncResponse{
		Type:      "node",
		Operation: "delete",
		OldObj:    obj,
	})
}

func (s *MetaDataWrapper) DeleteSerivce(obj interface{}) {
	kubernetes.DeleteService(obj)
	s.boardcast(api.MetaDataRsyncResponse{
		Type:      "service",
		Operation: "delete",
		OldObj:    obj,
	})
}

func (s *MetaDataWrapper) boardcast(data api.MetaDataRsyncResponse) {
	b, _ := json.Marshal(data)

	for _, flusher := range s.flushers {
		flusher.Write(append(b, '\n'))
	}
}

func (mp *MetaDataWrapper) list() ([]byte, error) {
	resp := api.ListVO{
		Cache:             kubernetes.MetaDataCache,
		GlobalNodeInfo:    kubernetes.GlobalNodeInfo,
		GlobalRsInfo:      kubernetes.GlobalRsInfo,
		GlobalServiceInfo: kubernetes.GlobalServiceInfo,
	}
	kubernetes.RLockMetadataCache()
	defer kubernetes.RUnlockMetadataCache()
	return json.Marshal(resp)
}

func GetIP(r *http.Request) (string, error) {
	ip := r.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	ip = r.Header.Get("X-Forward-For")
	for _, i := range strings.Split(ip, ",") {
		if net.ParseIP(i) != nil {
			return i, nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	return "", errors.New("no valid ip found")
}

func (mp *MetaDataWrapper) ListAndWatch(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	f := ioutil.NewWriteFlusher(w)
	//List Result
	b, _ := mp.list()
	f.Write(append(b, '\n'))
	//Watch Result
	ip, err := GetIP(r)
	if err == nil {
		log.Printf("Add Watcher , client is from IP: %s ,  current Watcher Size : %d", ip, len(mp.flushers))
	}
	mp.AddWatcher(ctx, f, 10, mp.stopCh)
}

type Watcher struct {
	eventChannel chan []byte
	*ioutil.WriteFlusher
}

func (mp *MetaDataWrapper) AddWatcher(ctx context.Context, flusher *ioutil.WriteFlusher, eventChannelSize int, stopCh <-chan struct{}) {
	w := &Watcher{
		eventChannel: make(chan []byte, eventChannelSize),
		WriteFlusher: flusher,
	}
	// TODO WaitGroup
	mp.flushers = append(mp.flushers, w)
	w.watch(ctx, stopCh)
	for index, watcher := range mp.flushers {
		if watcher == w {
			mp.flushers[index] = mp.flushers[len(mp.flushers)-1]
			mp.flushers = mp.flushers[:len(mp.flushers)-1]
		}
	}
}

func (w *Watcher) watch(ctx context.Context, stopCh <-chan struct{}) {
	for {
		select {
		case data := <-w.eventChannel:
			_, err := w.WriteFlusher.Write(data)
			if err != nil {
				// Remote Close this connection
				// TODO
				fmt.Println("Remote Connection Closed,Stop Watch!")
				return
			}
		case <-ctx.Done():
			return
		case <-stopCh:
			// TODO clear eventChannel and return
			return
		}
	}
}

func (mp *MetaDataWrapper) Shutdown() {
	// TODO Check channel
	close(mp.stopCh)
}
