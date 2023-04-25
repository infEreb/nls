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
NETFLOW_REPO_DIR=$REPO_DIR/netflow/server
NETFLOW_DIR=$BUILD_DIR/netflow
echo "Variable NETFLOW_REPO_DIR is $NETFLOW_REPO_DIR"
if [ ! -d $NETFLOW_DIR ]; then
    echo "Not found $NETFLOW_DIR directory, creating this one..."
    mkdir -p $NETFLOW_DIR/configs
fi
echo "Copying netflow data..."
cp $NETFLOW_REPO_DIR/Dockerfile $NETFLOW_DIR

NETFLOW_REPO_DIR=$NETFLOW_REPO_DIR/app
echo "Changing NETFLOW_REPO_DIR to $NETFLOW_REPO_DIR"
cd $NETFLOW_REPO_DIR

# Agent build
AGENT_DIR=$NETFLOW_DIR/agent
if [ ! -d $AGENT_DIR ]; then
    echo "Not found $AGENT_DIR directory, creating this one..."
    mkdir -p $AGENT_DIR/configs
fi
echo "Copying netflow/agent data..."
cp $NETFLOW_REPO_DIR/agent/configs/* $AGENT_DIR/configs
cp $NETFLOW_REPO_DIR/agent/Dockerfile $AGENT_DIR
go env -w GOBIN=$AGENT_DIR
go install ./agent/cmd/agent


# Ethernet build
ETHERNET_DIR=$NETFLOW_DIR/ethernet
if [ ! -d $ETHERNET_DIR ]; then
    echo "Not found $ETHERNET_DIR directory, creating this one..."
    mkdir -p $ETHERNET_DIR/configs
fi
echo "Copying netflow/ethernet data..."
cp $NETFLOW_REPO_DIR/ethernet/configs/* $ETHERNET_DIR/configs
cp $NETFLOW_REPO_DIR/ethernet/Dockerfile $ETHERNET_DIR
go env -w GOBIN=$ETHERNET_DIR
go install ./ethernet/cmd/ethernet
