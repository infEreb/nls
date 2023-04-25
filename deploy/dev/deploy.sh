#!/bin/bash
BUILD_VERSION=1.0.0

docker build -t nls-build ./nls-build
docker tag nls-build nls/nls-build:$BUILD_VERSION
#docker push nls/nls-build:$BUILD_VERSION