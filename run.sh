#!/bin/bash
#
# Title:run.sh
# Description:
#
PATH=/bin:/usr/bin:/etc:/usr/local/bin; export PATH
#
#export DB_HOST="192.168.1.102"
export DB_HOST="192.168.170.122"
export DB_HOST="localhost"
export DB_PASSWORD="goboy"
export DB_NAME="turtle_v1"
export DB_USER="turtle_go"
export FEATURE_FLAGS="3"
export GRPC_PORT="0.0.0.0:8080"
#192.168.170.122
#
./server
#
