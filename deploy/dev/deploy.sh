#!/bin/bash
BUILD_VERSION=1.0.0

docker build -t nls-build ./nls-build
docker tag nls-build nls/nls-build:$BUILD_VERSION
#docker push nls/nls-build:$BUILD_VERSION

docker run -d \
    -v build:/build
    nls-build:$BUILD_VERSION

# Netflow
docker build -t nls-netflow ./build/netflow
docker tag nls-netflow nls/nls-netflow:$BUILD_VERSION
# Netflow Agent
docker build -t nls-netflow-agent ./build/netflow/agent
docker tag nls-netflow-agent nls/nls-netflow-agent:$BUILD_VERSION
# Netflow Ethernet
docker build -t nls-netflow-ethernet ./build/netflow/ethernet
docker tag nls-netflow-ethernet nls/nls-netflow-ethernet:$BUILD_VERSION