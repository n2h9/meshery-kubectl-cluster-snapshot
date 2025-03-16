## yamls

yamls are generated from helm charts from meshery repo

```sh
helm --namespace meshery template meshery-operator install/kubernetes/helm/meshery-operator > meshery-operator.yaml
```

```sh
helm --namespace meshery template meshery-core install/kubernetes/helm/meshery > meshery-core.yaml
```

crds.yaml is copied from `install/kubernetes/helm/meshery-operator/crds`

For meshery core there are already k8s yamls in `install/deployment_yamls/k8s`.

For meshery operator not.

For now using generated yamls from helm.