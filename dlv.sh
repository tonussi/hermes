#!/bin/sh
cd ./cmd/http-log-server
dlv debug --headless --log -l 0.0.0.0:2345 --api-version=2
