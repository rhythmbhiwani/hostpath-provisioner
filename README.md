# HostPath Provisioner

A Kubernetes dynamic provisioning system for HostPath volumes, implemented in Go. This repository provides a Helm chart for deploying the provisioner and instructions for usage.

## Introduction

This project provides a dynamic provisioning system for Kubernetes using HostPath volumes. The HostPath provisioner dynamically creates HostPath persistent volumes on the nodes where the provisioner is running.

## Prerequisites

- Kubernetes cluster (v1.20+)
- Helm 3.0+

## Installation

```bash
helm install hostpath-provisioner \
  --namespace hostpath-provisioner \
  --create-namespace \
  https://github.com/rhythmbhiwani/hostpath-provisioner/releases/download/hostpath-provisioner-v1.0.4/hostpath-provisioner-v1.0.4.tgz
```

## Usage Instructions

### Deploy a PersistentVolumeClaim

Create a PersistentVolumeClaim (PVC) using the hostpath storage class:

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: my-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
  storageClassName: hostpath
```

Apply the PVC:

```bash
kubectl apply -f my-pvc.yaml
```

### Check the PersistentVolume

Verify that the PersistentVolume (PV) has been created and bound to your PVC:

```bash
kubectl get pv
kubectl get pvc
```
