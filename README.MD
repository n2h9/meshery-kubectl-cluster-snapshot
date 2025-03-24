## Build

run
```sh
./install-00.sh
```

copy the generted sha256 hash to meshery-cluster-snapshot.yaml sha256

## Install krew plugi

run
```sh
./install-01.sh
```

## Run plugin
```sh
kubectl meshery-cluster-snapshot --meshsyncBinaryPath=path_to_meshsync/meshsync --outputFile=meshery-cluster-snapshot-20250324-01.yaml
```

or  from binary
```sh
./bin/kubectl-meshery-cluster-snapshot --meshsyncBinaryPath=path_to_meshsync/meshsync --outputFile=meshery-cluster-snapshot-20250324-01.yaml
```

meshsyncBinaryPath is mandatory for now
binary must be build from this branch: https://github.com/meshery/meshsync/pull/386
(it contains output to file funcvtionality)