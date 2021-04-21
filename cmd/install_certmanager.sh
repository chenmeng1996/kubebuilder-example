#!/bin/sh

kubectl create namespace cert-manager
helm repo add jetstack https://charts.jetstack.io
helm repo update
# Note! https://cert-manager.io/docs/installation/kubernetes/ 如文档所示，v1.2.0版本开始支持的k8s版本>=1.16
# 1.14版本的k8s可以使用v1.1.0版本的cert-manager
helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --version v1.2.0 \
  --create-namespace \
  --set installCRDs=true