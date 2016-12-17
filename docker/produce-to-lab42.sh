#!/bin/sh

NETWORK=docker_default
CLIENT_CMD=/source/kafka-client
CLIENT_ARGS=$@

docker run \
  --net $NETWORK -v $(pwd)/../goclient:/source \
  -it pryz/confluent-goclient \
  $CLIENT_CMD $CLIENT_ARGS
