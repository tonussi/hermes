#!/usr/bin/env sh
ADDR=$1
PORT=$2
REQUESTS_N=$3

rm -rf /tmp/logs

for i in $(seq $(expr $REQUESTS_N))
do
    printf "\n"
    echo "writing on hermes"
    echo "{\"batch\":[{\"operation\":\"INSERT\",\"name\":\"name-$i\",\"city\":\"city-$i\"}]}"
    printf "\n"
    curl --request POST http://$ADDR:$PORT/db --data-raw "{\"batch\":[{\"operation\":\"INSERT\",\"name\":\"name-$i\",\"city\":\"city-$i\"}]}"
    printf "\n"
    echo "reading on hermes"
    printf "\n"
    curl --request POST http://$ADDR:$PORT/line --data-raw "{\"number\":-1}"
    printf "\n"
done

for i in $(seq $(expr $REQUESTS_N + 2))
do
    printf "\n"
    echo "reading on hermes"
    printf "\n"
    curl --request POST http://$ADDR:$PORT/line --data-raw "{\"number\":$i}"
    printf "\n"
done

wc -l /tmp/logs/operations.log

cat /tmp/logs/operations.log
