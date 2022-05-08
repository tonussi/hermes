#!/bin/sh
docker build -t tonussi/studygo-http-log-server -f dockers/http-log-server/debug.dockerfile .
docker run -p 5000:5000 --name studygo-http-log-server tonussi/studygo-http-log-server
