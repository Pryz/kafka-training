#!/bin/sh

#
# Create lab42 topic with 1 partition
#

DOCKER_IMG=wurstmeister/kafka
NETWORK=docker_default
KAFKA_CMD='/opt/kafka/bin/kafka-topics.sh --zookeeper zk:2181 --create --replication-factor 1 --partitions 1 --topic lab42'

docker run --net $NETWORK $DOCKER_IMG $KAFKA_CMD
