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

start_minikube:
	cd infra; ./start_minikube.sh

redis_deploy:
	$(KUBECTL) create namespace cyclops-redis	
	cd infra; $(KUBECTL) apply -f redis-secret.yaml -n cyclops-redis
	cd infra; $(HELM) install cyclops-redis bitnami/redis --values redis-values.yaml -n cyclops-redis

manager_build:
	cd manager; $(DOCKER) build . -t $(DARING_CYCLOPS_MANAGER)

manager_deploy:
	$(KUBECTL) create namespace cyclops-app	
	cd manager; $(KUBECTL) apply -f manager-deploy.yaml

worker_build:
	cd worker; $(DOCKER) build . -t $(DARING_CYCLOPS_WORKER)

worker_deploy:
	$(KUBECTL) create namespace cyclops-app
	cd infra; $(KUBECTL) apply -f worker-deploy.yaml -n cyclops-app

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
