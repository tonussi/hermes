#!/usr/bin/env sh
DOCKERHUB_USER_NAME=lptonussi

docker build -t $DOCKERHUB_USER_NAME/public:hermes -f docker/hermes.dockerfile .
# docker build -t $DOCKERHUB_USER_NAME/public:tcp-kv-client -f docker/tcp-kv-client.dockerfile .
# docker build -t $DOCKERHUB_USER_NAME/public:tcp-kv-hashicorp-raft -f docker/tcp-kv-hashicorp-raft.dockerfile .
# docker build -t $DOCKERHUB_USER_NAME/public:tcp-kv-server -f docker/tcp-kv-server.dockerfile .

docker push $DOCKERHUB_USER_NAME/public:hermes
# docker push $DOCKERHUB_USER_NAME/public:tcp-kv-client
# docker push $DOCKERHUB_USER_NAME/public:tcp-kv-hashicorp-raft
# docker push $DOCKERHUB_USER_NAME/public:tcp-kv-server

# docker build -t $DOCKERHUB_USER_NAME/public:golang-http-log-server -f client.dockerfile .
# docker build -t $DOCKERHUB_USER_NAME/public:golang-http-log-client -f server.dockerfile .

# docker push $DOCKERHUB_USER_NAME/public:golang-http-log-server
# docker push $DOCKERHUB_USER_NAME/public:golang-http-log-client
