#!/bin/bash
#
# Title:start_ingress.sh
# Description: start minikube ingress
# Development Environment: OS X 10.13.6
# Author: G.S. Cole (guycole at gmail dot com)
#
minikube start --memory=6g --bootstrapper=kubeadm --extra-config=kubelet.authentication-token-webhook=true --extra-config=kubelet.authorization-mode=Webhook --extra-config=scheduler.address=0.0.0.0 --extra-config=controller-manager.address=0.0.0.0 --vm=true
minikube addons enable ingress
#
