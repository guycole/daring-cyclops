#!/bin/bash
#
# Title:server.sh
# Description:
#
PATH=/bin:/usr/bin:/etc:/usr/local/bin; export PATH
#
export FEATURE_FLAGS="1"
export GRPC_ADDRESS="0.0.0.0:8080"
export RUN_MODE="server"
#
./daring-cyclops
#
