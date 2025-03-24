#!/bin/bash

echo "WARNING: this will update checksum, and if you want to install krew plugin you need to update meshery-cluster-snapshot.yaml sha256 field"

kubectl krew uninstall meshery-cluster-snapshot 2>/dev/null

go build -o bin/kubectl-meshery-cluster-snapshot cmd/meshsync/*.go

tar -czvf bin/v0.0.1.tar.gz bin/kubectl-meshery-cluster-snapshot

sha256sum bin/v0.0.1.tar.gz