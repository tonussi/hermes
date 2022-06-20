#!/usr/bin/env sh
DOCKERHUB_USER_NAME=lptonussi

docker build -t $DOCKERHUB_USER_NAME/hermes -f docker/hermes.dockerfile .

docker push $DOCKERHUB_USER_NAME/hermes
