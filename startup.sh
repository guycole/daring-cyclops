#!/bin/bash
#
# Title:startup.sh
# Description: 
# Development Environment: OS X 10.13.6
# Author: G.S. Cole (guycole at gmail dot com)
#
make minikube_start
make minikube_setup
#
make redis_deploy
#
make manager_build
make worker_build
#
make worker_apply
make manager_apply
#
make monitoring_deploy
#