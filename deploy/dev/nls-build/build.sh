#!/bin/bash

REPO_DIR=/nls
BUILD_DIR=/build

# /build/
# /build/service/
# /build/service/app
# /build/service/config/conf
if [ ! -d $REPO_DIR ]; then
    mkdir  $BUILD_DIR
fi


git clone https://github.com/infEreb/nls.git $REPO_DIR
# Get NETFLOW repo
NETFLOW_REPO_DIR=$REPO_DIR/netflow/server/app
NETFLOW_DIR=$BUILD_DIR/netflow
if [ ! -d $NETFLOW_DIR ]; then
    mkdir -p $NETFLOW_DIR/configs
fi
cp $NETFLOW_REPO_DIR/netflow/Dockerfile $NETFLOW_DIR

# Agent build
AGENT_DIR=$NETFLOW_DIR/agent
if [ ! -d $AGENT_DIR ]; then
    mkdir -p $AGENT_DIR/configs
fi

cp $NETFLOW_REPO_DIR/agent/configs/* $AGENT_DIR/configs
cp $NETFLOW_REPO_DIR/agent/Dockerfile $AGENT_DIR
go env GOBIN=$AGENT_DIR
go install $AGENT_DIR/cmd/agent


# Ethernet build
ETHERNET_DIR=$NETFLOW_DIR/ethernet
if [ ! -d $ETHERNET_DIR ]; then
    mkdir -p $ETHERNET_DIR/configs
fi

cp $NETFLOW_REPO_DIR/ethernet/configs/* $ETHERNET_DIR/configs
cp $NETFLOW_REPO_DIR/ethernet/Dockerfile $ETHERNET_DIR
go env GOBIN=$ETHERNET_DIR
go install $ETHERNET_DIR/cmd/ethernet
