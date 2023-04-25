BUILD_VERSION=1.0.0

docker run -d -p 5000:5000 --restart=always --name registry registry:2

docker build -t nls-build:$BUILD_VERSION ./nls-build
docker tag nls-build:$BUILD_VERSION localhost:5000/nls-build:$BUILD_VERSION
docker push localhost:5000/nls-build:$BUILD_VERSION

docker run -d \
    -v $(pwd)/build:/build \
    --name build \
    nls-build:$BUILD_VERSION

# Netflow
docker build -t nls-netflow:$BUILD_VERSION ./build/netflow
docker tag nls-netflow:$BUILD_VERSION localhost:5000/nls-netflow:$BUILD_VERSION
docker push localhost:5000/nls-netflow:$BUILD_VERSION
# Netflow Agent
docker build -t nls-netflow-agent:$BUILD_VERSION ./build/netflow/agent
docker tag nls-netflow-agent:$BUILD_VERSION localhost:5000/nls-netflow-agent:$BUILD_VERSION
docker push localhost:5000/nls-netflow-agent:$BUILD_VERSION
# Netflow Ethernet
docker build -t nls-netflow-ethernet:$BUILD_VERSION ./build/netflow/ethernet
docker tag nls-netflow-ethernet:$BUILD_VERSION localhost:5000/nls-netflow-ethernet:$BUILD_VERSION
docker push localhost:5000/nls-netflow-ethernet:$BUILD_VERSION