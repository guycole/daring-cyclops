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

manager_build:
	cd manager; $(DOCKER) build . -t $(DARING_CYCLOPS_MANAGER)

manager_apply:
	cd infra; $(KUBECTL) apply -f manager-deploy.yaml

manager_delete:
	cd infra; $(KUBECTL) delete -f manager-deploy.yaml -n cyclops-app

minikube_start:
	cd infra; ./start_minikube.sh

minikube_setup:
	$(KUBECTL) create namespace cyclops-app	
	$(KUBECTL) create namespace monitoring

monitoring_deploy:
	cd infra; $(HELM) install prometheus prometheus-community/kube-prometheus-stack --namespace monitoroing  

redis_deploy:
	cd infra; $(KUBECTL) apply -f redis-secret.yaml -n cyclops-app
	cd infra; $(HELM) install cyclops-redis bitnami/redis --values redis-values.yaml -n cyclops-app

worker_build:
	cd worker; $(DOCKER) build . -t $(DARING_CYCLOPS_WORKER)

worker_apply:
	cd infra; $(KUBECTL) apply -f worker-deploy.yaml -n cyclops-app

worker_delete:
	cd infra; $(KUBECTL) delete -f worker-deploy.yaml -n cyclops-app

#
#  Cleanup this subdirectory.
#
.PHONY: clean
clean:
	-@rm -f *.o *.BAK core

#
#  Nuke all the executables.
#
.PHONY ultraclean:
ultraclean:
	-@rm -f lib/*.a
	-@rm -f *~ TAGS depend.include $(TEST1)
	-@touch depend.include
