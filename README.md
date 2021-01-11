# Redpanda operator

[Redpanda operator](https://github.com/vectorizedio/kubernetes-operator) is Kubernetes Operator
for managing and automating tasks related to managing Redpanda clusters.

[Redpanda](https://github.com/vectorizedio/redpanda) is a streaming platform for mission critical
workloads. Kafka® compatible, No Zookeeper®, no JVM, and no code changes required.
Use all your favorite open source tooling - 10x faster.

## Getting started

### Requirements

* Kubernetes 1.16 or newer
* kubectl 1.16 or newer
* kustomize v3.8.7 or newer
* kind v0.9.0 or newer

### Installation

#### Local installation

First clone the repo

```
git clone https://github.com/vectorizedio/kubernetes-operator.git
```

Create local Kubernetes cluster using KIND

```
kind create cluster -config kind.yaml
```

Build and load the container to Kubernetes cluster

```
make docker-build
make push-to-kind
```

Then build the Kubernetes manifest and apply them to cluster

```
kustomize build config/default | kubectl apply -f -
```

Install sample RedpandaCluster custom resource

```
kubectl apply -f config/samples/core_v1alpha1_redpandacluster.yaml
```
