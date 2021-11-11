#!/bin/bash
#
# Title:start_minikube.sh
# Description: start minikube
# Development Environment: OS X 10.13.6
# Author: G.S. Cole (guycole at gmail dot com)
#
minikube start --memory=6g --nodes=1 --bootstrapper=kubeadm --extra-config=kubelet.authentication-token-webhook=true --extra-config=kubelet.authorization-mode=Webhook --extra-config=scheduler.address=0.0.0.0 --extra-config=controller-manager.address=0.0.0.0 --vm=true --kubernetes-version=v1.20.0
minikube addons enable ingress
#