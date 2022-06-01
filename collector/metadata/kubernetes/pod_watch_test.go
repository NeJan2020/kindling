package kubernetes

import (
	"testing"

	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestTruncateContainerId(t *testing.T) {
	testCases := []struct {
		containerId string
		expected    string
	}{
		{"docker://a1b2c3d4e5f6g7h8i9j0k1l2m3n", "a1b2c3d4e5f6"},
		{"docker://a1b2c3d4e5f6", "a1b2c3d4e5f6"},
		{"docker://a1b2c3", "a1b2c3"},
		{"containerd://a1b2c3d4e5f6g7h8i9j0k1l2m3n", "a1b2c3d4e5f6"},
		{"a1b2c3", ""},
		{"", ""},
	}
	for _, test := range testCases {
		res := TruncateContainerId(test.containerId)
		if res != test.expected {
			t.Errorf("containerId=%s, get=%s, but expected=%s", test.containerId, res, test.expected)
		}
	}
}

func TestOnAdd(t *testing.T) {
	globalPodInfo = &podMap{
		Info: make(map[string]map[string]*PodInfo),
	}
	globalServiceInfo = &ServiceMap{
		ServiceMap: make(map[string]map[string]*K8sServiceInfo),
	}
	globalRsInfo = &ReplicaSetMap{
		Info: make(map[string]Controller),
	}
	// First add service, and then add pod
	onAddService(CreateService())
	onAddReplicaSet(CreateReplicaSet())
	onAdd(CreatePod(true))
	t.Log(MetaDataCache)
	// Delete podInfo must not affect serviceMap
	onDelete(CreatePod(true))
	t.Log(MetaDataCache)
	// Empty all the metadata
	onDeleteService(CreateService())
	t.Log(MetaDataCache)
}

// ISSUE https://github.com/CloudDectective-Harmonycloud/kindling/issues/229
func TestOnAddPodWhileReplicaSetUpdating(t *testing.T) {
	globalPodInfo = &podMap{
		Info: make(map[string]map[string]*PodInfo),
	}
	globalServiceInfo = &ServiceMap{
		ServiceMap: make(map[string]map[string]*K8sServiceInfo),
	}
	globalRsInfo = &ReplicaSetMap{
		Info: make(map[string]Controller),
	}
	// Firstly deployment created and add old RS and old POD
	controller := true
	oldRs := CreateReplicaSet()
	oldRs.SetResourceVersion("old")
	newRs := CreateReplicaSet()
	newRs.SetResourceVersion("new")
	oldPOD := CreatePod(true)
	oldPOD.SetResourceVersion("old")
	oldPOD.OwnerReferences[0].Controller = &controller
	newPOD := CreatePod(true)
	newPOD.SetResourceVersion("new")
	newPOD.OwnerReferences[0].Controller = &controller
	onAddReplicaSet(oldRs)
	onAdd(oldPOD)

	// Secondly POD&RS were been updated

	go func() {
		for i := 0; i < 1000; i++ {
			OnUpdateReplicaSet(oldRs, newRs)
		}
	}()

	for i := 0; i < 100; i++ {
		OnUpdate(oldPOD, newPOD)
		// Thirdly check the pod's workload_kind
		pod, ok := MetaDataCache.GetPodByContainerId(TruncateContainerId(newPOD.Status.ContainerStatuses[0].ContainerID))
		require.True(t, ok, "failed to get target POD")
		require.Equal(t, "deployment", pod.WorkloadKind, "failed to get the real workload_kind")
	}
}

func TestOnAddLowercaseWorkload(t *testing.T) {
	globalPodInfo = &podMap{
		Info: make(map[string]map[string]*PodInfo),
	}
	globalServiceInfo = &ServiceMap{
		ServiceMap: make(map[string]map[string]*K8sServiceInfo),
	}
	globalRsInfo = &ReplicaSetMap{
		Info: make(map[string]Controller),
	}
	higherCase := "DaemonSet"
	lowerCase := "daemonset"
	isController := true
	onAdd(&corev1.Pod{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			OwnerReferences: []metav1.OwnerReference{{
				Kind:       higherCase,
				Controller: &isController,
			}},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name: "container1",
					Ports: []corev1.ContainerPort{
						{
							ContainerPort: 80,
						},
					},
				},
			}},
		Status: corev1.PodStatus{
			PodIP: "172.10.1.2",
			ContainerStatuses: []corev1.ContainerStatus{
				{
					Name:        "container1",
					ContainerID: "docker://1a2b3c4d5e6f7g8h9i0j1k2",
				},
			},
		},
	})
	podInfo, ok := MetaDataCache.GetPodByContainerId("1a2b3c4d5e6f")
	if !ok || podInfo.WorkloadKind != lowerCase {
		t.Errorf("%s wanted, but get %s", higherCase, lowerCase)
	}
}

func CreatePod(hasPort bool) *corev1.Pod {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "deploy-1a2b3c4d-5e6f7",
			Namespace: "CustomNamespace",
			Labels: map[string]string{
				"a": "1",
				"b": "1",
			},
			OwnerReferences: []metav1.OwnerReference{
				{
					Kind: ReplicaSetKind,
					Name: "deploy-1a2b3c4d",
				},
			},
		},
		Spec: corev1.PodSpec{
			NodeName:    "node1",
			HostNetwork: false,
		},
		Status: corev1.PodStatus{
			PodIP: "172.10.1.2",
			ContainerStatuses: []corev1.ContainerStatus{
				{
					Name:        "container1",
					ContainerID: "docker://1a2b3c4d5e6f7g8h9i0j1k2",
				},
			},
		},
	}
	if hasPort {
		pod.Spec.Containers = []corev1.Container{
			{
				Name: "container1",
				Ports: []corev1.ContainerPort{
					{
						ContainerPort: 80,
					},
				},
			},
		}
	} else {
		pod.Spec.Containers = []corev1.Container{
			{
				Name: "container1",
			},
		}
	}
	return pod
}
