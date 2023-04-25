# Build container

FROM golang:1.19

RUN apt-get install -y bash wget tar git

COPY build.sh .

CMD ./build.sh

# docker run nls-build -v local/build:/build -v 