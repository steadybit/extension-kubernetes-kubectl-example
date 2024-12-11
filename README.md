# Copy & Paste from extension-kubernetes

Showcase to run a kubectl command in a selected Kubernetes-Container


## Installation / De-Installation

```bash
helm repo add extension-kubernetes-kubectl-example https://steadybit.github.io/extension-kubernetes-kubectl-example
helm repo update

helm install extension-kubernetes-kubectl-example --namespace steadybit-agent --set image.tag=main --set image.pullPolicy=Always --set kubernetes.clusterName=MY-CLUSTER extension-kubernetes-kubectl-example/steadybit-extension-kubernetes-kubectl-example

helm uninstall extension-kubernetes-kubectl-example --namespace steadybit-agent
```
