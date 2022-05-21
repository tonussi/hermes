#!/usr/bin/env sh
DOCKERHUB_USER_NAME=lptonussi

docker build -t $DOCKERHUB_USER_NAME/public:golang-http-log-server -f client.dockerfile .
docker build -t $DOCKERHUB_USER_NAME/public:golang-http-log-client -f server.dockerfile .

docker push $DOCKERHUB_USER_NAME/public:golang-http-log-server
docker push $DOCKERHUB_USER_NAME/public:golang-http-log-client
