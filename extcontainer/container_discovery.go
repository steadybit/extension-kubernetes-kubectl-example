// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2023 Steadybit GmbH

package extcontainer

import (
	"context"
	"fmt"
	"github.com/steadybit/discovery-kit/go/discovery_kit_api"
	"github.com/steadybit/discovery-kit/go/discovery_kit_sdk"
	"github.com/steadybit/extension-kit/extutil"
	"github.com/steadybit/extension-kubernetes-kubectl-example/client"
	"github.com/steadybit/extension-kubernetes-kubectl-example/extcommon"
	"github.com/steadybit/extension-kubernetes-kubectl-example/extconfig"
	"github.com/steadybit/extension-kubernetes-kubectl-example/extnamespace"
	corev1 "k8s.io/api/core/v1"
	"reflect"
	"strings"
	"time"
)

type containerDiscovery struct {
	k8s *client.Client
}

var (
	_ discovery_kit_sdk.Discovery       = (*containerDiscovery)(nil)
	_ discovery_kit_sdk.TargetDiscovery = (*containerDiscovery)(nil)
)

const KubernetesContainerTargetType = "com.steadybit.extension_kubernetes-kubectl-example.kubernetes-container-demo"

func NewContainerDiscovery(ctx context.Context, k8s *client.Client) discovery_kit_sdk.TargetDiscovery {
	discovery := &containerDiscovery{k8s: k8s}
	chRefresh := extcommon.TriggerOnKubernetesResourceChange(k8s, reflect.TypeOf(corev1.Pod{}), reflect.TypeOf(corev1.Node{}))
	return discovery_kit_sdk.NewCachedTargetDiscovery(
		discovery,
		discovery_kit_sdk.WithRefreshTargetsNow(),
		discovery_kit_sdk.WithRefreshTargetsTrigger(ctx, chRefresh, 5*time.Second),
	)
}

func (c *containerDiscovery) Describe() discovery_kit_api.DiscoveryDescription {
	return discovery_kit_api.DiscoveryDescription{
		Id: KubernetesContainerTargetType,
		Discover: discovery_kit_api.DescribingEndpointReferenceWithCallInterval{
			CallInterval: extutil.Ptr("30s"),
		},
	}
}

func (c *containerDiscovery) DiscoverTargets(_ context.Context) ([]discovery_kit_api.Target, error) {
	pods := c.k8s.Pods()

	targetList := make([]discovery_kit_api.Target, 0, len(pods))
	for _, pod := range pods {
		podMetadata := pod.ObjectMeta
		ownerReferences := client.OwnerReferences(c.k8s, &podMetadata)
		services := c.k8s.ServicesMatchingToPodLabels(pod.Namespace, pod.ObjectMeta.Labels)

		for _, container := range pod.Status.ContainerStatuses {
			if container.ContainerID == "" {
				continue
			}
			containerIdWithoutPrefix := strings.SplitAfter(container.ContainerID, "://")[1]

			attributes := map[string][]string{
				"k8s.cluster-name":          {extconfig.Config.ClusterName},
				"k8s.container.id":          {container.ContainerID},
				"k8s.container.id.stripped": {containerIdWithoutPrefix},
				"k8s.container.name":        {container.Name},
				"k8s.container.image":       {container.Image},
				"k8s.namespace":             {podMetadata.Namespace},
				"k8s.node.name":             {pod.Spec.NodeName},
				"k8s.pod.name":              {podMetadata.Name},
				"k8s.distribution":          {c.k8s.Distribution},
			}

			for key, value := range podMetadata.Labels {
				attributes[fmt.Sprintf("k8s.pod.label.%v", key)] = []string{value}
				attributes[fmt.Sprintf("k8s.label.%v", key)] = []string{value}
			}
			extnamespace.AddNamespaceLabels(c.k8s, pod.Namespace, attributes)
			extcommon.AddNodeLabels(c.k8s.Nodes(), pod.Spec.NodeName, attributes)

			if len(services) > 0 {
				var serviceNames = make([]string, 0, len(services))
				for _, service := range services {
					serviceNames = append(serviceNames, service.Name)
				}
				attributes["k8s.service.name"] = serviceNames
			}

			for _, ownerRef := range ownerReferences.OwnerRefs {
				attributes[fmt.Sprintf("k8s.%v", ownerRef.Kind)] = []string{ownerRef.Name}
				attributes["k8s.workload-type"] = []string{ownerRef.Kind}
				attributes["k8s.workload-owner"] = []string{ownerRef.Name}
			}

			targetList = append(targetList, discovery_kit_api.Target{
				Id:         container.ContainerID,
				Label:      container.Name,
				TargetType: KubernetesContainerTargetType,
				Attributes: attributes,
			})
		}
	}
	return targetList, nil
}
