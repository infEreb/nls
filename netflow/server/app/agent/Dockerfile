FROM alpine:latest

RUN apk add -y bash

WORKDIR /app
COPY ./configs .
COPY ./build/. .

CMD ./agent-service -c /app/configs/conf.yaml