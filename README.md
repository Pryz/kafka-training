# Kafka training

## Start cluster

Start Zookeeper and 1 Kafka broker

```
docker-compose up
```

## "Connect" to Kafka

```
docker exec -it docker_kafka_1

ps auxwww | grep kafka
```

## Create topic

```
/opt/kafka/bin/kafka-topics.sh --zookeeper zk:2181 --create --replication-factor 1 --partitions 1 --topic lab42
```

## Produce to topic

```
echo "helloworld" | \
/opt/kafka/bin/kafka-console-producer.sh --broker-list localhost:9092 --topic lab42
```

## Consume to topic

```
/opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic lab42 --from-beginning
```

## Audit cluster (list topics, leaders, ...)

```
/opt/kafka/bin/kafka-topics.sh --zookeeper zk:2181 --describe
```

## Add two brokers in the cluster

```
docker-compose scale kafka=3
```

## Look at the current topic repartition

```
docker exec -it docker_kafka_1

/opt/kafka/bin/kafka-topics.sh --zookeeper zk:2181 --describe --topic lab42
```

## Change leader

```
```

## Edit topic

```
```

