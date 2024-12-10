package extnamespace

import (
	"fmt"
	"github.com/steadybit/extension-kubernetes-kubectl-example/client"
)

func AddNamespaceLabels(k8s *client.Client, namespace string, attributes map[string][]string) map[string][]string {
	if k8s.Permissions().CanReadNamespaces() {
		namespaces := k8s.Namespaces()
		for _, ns := range namespaces {
			if ns.Name == namespace {
				namespaceMetadata := ns.ObjectMeta
				for key, value := range namespaceMetadata.Labels {
					attributes[fmt.Sprintf("k8s.namespace.label.%v", key)] = []string{value}
					attributes[fmt.Sprintf("k8s.label.%v", key)] = []string{value}
				}
				break
			}
		}
	}
	return attributes
}
