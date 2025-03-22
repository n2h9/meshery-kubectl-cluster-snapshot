## commands

```sh
go build -o bin/kubectl-meshery-cluster-snapshot cmd/meshkit/*.go
```

```sh
tar -czvf bin/v0.0.1.tar.gz bin/kubectl-meshery-cluster-snapshot
```

```sh
sha256sum bin/v0.0.1.tar.gz
```

```sh
kubectl krew install --manifest=meshery-cluster-snapshot.yaml --archive=bin/v0.0.1.tar.gz
```

```sh
kubectl meshery-cluster-snapshot
```

```sh
kubectl krew uninstall meshery-cluster-snapshot
```