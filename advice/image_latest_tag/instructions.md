Makes sure that your Kubernetes manifest is based on an ```image``` with a fixed and explicit tag

```yaml
apiVersion: v1
kind: Pod
spec:
  containers:
    - name:  ${target.k8s.deployment}
% startHighlight %
      image: images.my-company.example/app:v4
% endHighlight %

```