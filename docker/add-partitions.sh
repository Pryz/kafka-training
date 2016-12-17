#!/bin/sh

#
# Add 2 partitions to lab42 topic
#

/opt/kafka/bin/kafka-topics.sh --zookeeper zk:2181 --alter --topic lab42 --partitions 3
