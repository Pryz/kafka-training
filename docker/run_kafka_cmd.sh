#!/bin/sh

#
# Execute Kafka command
#

DOCKER_IMG=wurstmeister/kafka
NETWORK=docker_default
KAFKA_CMD=$1

docker run --net $NETWORK $DOCKER_IMG $KAFKA_CMD
