// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2023 Steadybit GmbH

package extdaemonset

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/steadybit/extension-kit/extutil"
	"github.com/steadybit/extension-kubernetes/client"
	"github.com/steadybit/extension-kubernetes/extconfig"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	testclient "k8s.io/client-go/kubernetes/fake"
	"testing"
	"time"
)

func Test_getDiscoveredDaemonSets(t *testing.T) {
	// Given
	stopCh := make(chan struct{})
	defer close(stopCh)
	client, clientset := getTestClient(stopCh)
	extconfig.Config.ClusterName = "development"
	extconfig.Config.LabelFilter = []string{"secret-label"}
	extconfig.Config.DiscoveryMaxPodCount = 50

	_, err := clientset.CoreV1().
		Pods("default").
		Create(context.Background(), &v1.Pod{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Pod",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "shop-pod",
				Namespace: "default",
				Labels: map[string]string{
					"best-city": "kevelaer",
				},
			},
			Status: v1.PodStatus{
				ContainerStatuses: []v1.ContainerStatus{
					{
						ContainerID: "crio://abcdef",
						Name:        "MrFancyPants",
						Image:       "nginx",
					},
				},
			},
			Spec: v1.PodSpec{
				NodeName: "worker-1",
				Containers: []v1.Container{
					{
						Name:            "nginx",
						Image:           "nginx",
						ImagePullPolicy: "Always",
					},
				},
			},
		}, metav1.CreateOptions{})
	require.NoError(t, err)

	sts, err := clientset.
		AppsV1().
		DaemonSets("default").
		Create(context.Background(), &appsv1.DaemonSet{
			TypeMeta: metav1.TypeMeta{
				Kind:       "DaemonSet",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "shop",
				Namespace: "default",
				Labels: map[string]string{
					"best-city":    "Kevelaer",
					"secret-label": "secret-value",
				},
			},
			Spec: appsv1.DaemonSetSpec{
				Selector: extutil.Ptr(metav1.LabelSelector{
					MatchLabels: map[string]string{
						"best-city": "kevelaer",
					},
				}),
			},
		}, metav1.CreateOptions{})
	require.NoError(t, err)
	log.Printf("Created daemonset %v", sts.Name)

	// When
	assert.Eventually(t, func() bool {
		return len(getDiscoveredDaemonSetTargets(client)) == 1
	}, time.Second, 100*time.Millisecond)

	// Then
	targets := getDiscoveredDaemonSetTargets(client)
	require.Len(t, targets, 1)
	target := targets[0]
	assert.Equal(t, "development/default/shop", target.Id)
	assert.Equal(t, "shop", target.Label)
	assert.Equal(t, DaemonSetTargetType, target.TargetType)
	assert.Equal(t, map[string][]string{
		"host.hostname":                                {"worker-1"},
		"k8s.namespace":                                {"default"},
		"k8s.daemonset":                                {"shop"},
		"k8s.label.best-city":                          {"Kevelaer"},
		"k8s.cluster-name":                             {"development"},
		"k8s.pod.name":                                 {"shop-pod"},
		"k8s.container.id":                             {"crio://abcdef"},
		"k8s.container.id.stripped":                    {"abcdef"},
		"k8s.distribution":                             {"kubernetes"},
		"k8s.container.spec.name.limit.cpu.not-set":    {"nginx"},
		"k8s.container.spec.name.limit.memory.not-set": {"nginx"},
		"k8s.container.probes.liveness.not-set":        {"nginx"},
		"k8s.container.probes.readiness.not-set":       {"nginx"},
	}, target.Attributes)
}

func getTestClient(stopCh <-chan struct{}) (*client.Client, kubernetes.Interface) {
	clientset := testclient.NewSimpleClientset()
	client := client.CreateClient(clientset, stopCh, "", client.MockAllPermitted())
	return client, clientset
}
