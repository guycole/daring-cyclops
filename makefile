#
# Title:makefile
#
# Description:
#
#   'make clean' removes all core and object files
#   'make ultraclean' removes all executables
#
# Operating System/Environment:
#   Ubuntu 18.04.3 LTS (Bionic Beaver)
#   GNU Make 4.1
#
# Author:
#   G.S. Cole (guycole at gmail dot com)
#
DOCKER = docker
DARING_CYCLOPS_MANAGER = daring-cyclops-manager:1
DARING_CYCLOPS_WORKER = daring-cyclops-worker:1
HELM = helm
KUBECTL = kubectl
MINIKUBE = minikube

manager_build:
	cd manager; $(DOCKER) build . -t $(DARING_CYCLOPS_MANAGER)

manager_delete:
	$(KUBECTL) delete -f infra/manager-ingress.yaml -n cyclops-app
	$(KUBECTL) delete -f infra/manager-service.yaml -n cyclops-app
	$(KUBECTL) delete -f infra/manager-deploy.yaml -n cyclops-app

manager_deploy:
	$(KUBECTL) apply -f infra/manager-deploy.yaml -n cyclops-app
	$(KUBECTL) apply -f infra/manager-service.yaml -n cyclops-app	
	$(KUBECTL) apply -f infra/manager-ingress.yaml -n cyclops-app

minikube_reset:
	$(MINIKUBE) stop
	$(MINIKUBE) delete

minikube_start:
	cd infra; ./start_minikube.sh

minikube_setup:
	$(KUBECTL) create namespace cyclops-app	
	$(KUBECTL) create namespace monitoring
	$(MINIKUBE) addons enable ingress
	$(HELM) repo add stable https://charts.helm.sh/stable
	$(HELM) repo update

monitoring_delete:
	$(KUBECTL) delete -f infra/redis-dashboard.yaml -n monitoring
	$(HELM) uninstall prometheus -n monitoring

monitoring_deploy:
	$(KUBECTL) apply -f infra/redis-dashboard.yaml -n monitoring
	$(HELM) repo add prometheus-community https://prometheus-community.github.io/helm-charts
	$(HELM) upgrade --debug --install prometheus prometheus-community/kube-prometheus-stack -n monitoring --version 19.0.2 --values infra/kube-prometheus.yaml

monitoring_expose:
	$(KUBECTL) expose service prometheus-kube-prometheus-alertmanager --type=NodePort --target-port=9093 --name=prometheus-alertmanager-np --namespace=monitoring
	$(KUBECTL) expose service prometheus-kube-prometheus-prometheus --type=NodePort --target-port=9090 --name=prometheus-np --namespace=monitoring
	$(KUBECTL) expose service prometheus-grafana --type=NodePort --target-port=3000 --name=grafana-np --namespace=monitoring

redis_deploy:
	$(KUBECTL) apply -f infra/redis-secret.yaml -n cyclops-app 
	$(HELM) repo add bitnami https://charts.bitnami.com/bitnami
	$(HELM) upgrade --debug --install cyclops-redis bitnami/redis -n cyclops-app --version 15.5.4 --values infra/redis-values.yaml

worker_build:
	cd worker; $(DOCKER) build . -t $(DARING_CYCLOPS_WORKER)

worker_delete:
	$(KUBECTL) delete -f infra/worker-deploy.yaml -n cyclops-app

worker_deploy:
	$(KUBECTL) apply -f infra/worker-deploy.yaml -n cyclops-app
