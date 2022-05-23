#!/usr/bin/env sh
curl --request POST http://www.localhost:8001/db --data-raw '{"batch":[{"operation":"INSERT","name":"luquas","city":"floripa"}]}'
curl --request POST http://www.localhost:8001/line --data-raw '{"number":2}'
