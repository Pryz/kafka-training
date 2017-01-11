# Kafka training

The purpose of this training is to give you a starting point to use and experiment with Kafka.

## Resources

Some interesting resources :

* The official Kafka documentation : [https://kafka.apache.org/documentation/](https://kafka.apache.org/documentation/)
* A training Deck of Kafka 0.8 : [http://www.slideshare.net/miguno/apache-kafka-08-basic-training-verisign](http://www.slideshare.net/miguno/apache-kafka-08-basic-training-verisign)
* Monitoring Kafka from Datadog folks : [https://www.datadoghq.com/blog/monitoring-kafka-performance-metrics/](https://www.datadoghq.com/blog/monitoring-kafka-performance-metrics/)
* Interesting reading for a Production deployment from Confluent : [http://docs.confluent.io/3.1.1/schema-registry/docs/deployment.html?highlight=production](http://docs.confluent.io/3.1.1/schema-registry/docs/deployment.html?highlight=production)

### Start cluster

Start containers :

```
docker-compose up
```

This will start :

* A Zookeeper node listening on port 2181
* A Kafka broker
* A Zookeeper web app to browse the ZK content listening on 4550. ([http://localhost:4550](http://localhost:4550)) from your web browser.

The Input connection string for the ZK browser will be : zk:2181.

### "Connect" to Kafka

```
docker exec -it docker_kafka_1 bash

ps auxwww | grep kafka
```

All the command perform against Kafka will happen in this container.

### Create lab42 topic

```
/opt/kafka/bin/kafka-topics.sh --zookeeper zk:2181 --create --replication-factor 1 --partitions 1 --topic lab42
```

### Produce to lab42 topic

```
echo "helloworld" | \
/opt/kafka/bin/kafka-console-producer.sh --broker-list localhost:9092 --topic lab42
```

### Consume from lab42 topic

```
/opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic lab42 --from-beginning
```

### Audit cluster (list topics, leaders, ...)

```
/opt/kafka/bin/kafka-topics.sh --zookeeper zk:2181 --describe
```

### Add two brokers in the cluster

```
docker-compose scale kafka=3
```

### Add partitions to lab42 topic

```
/opt/kafka/bin/kafka-topics.sh --zookeeper zk:2181 --alter --topic lab42 --partitions 3
```

### Look at the current topic repartition

```
docker exec -it docker_kafka_1

/opt/kafka/bin/kafka-topics.sh --zookeeper zk:2181 --describe --topic lab42
```

### Change leader

Create a file like the following :

```
$ cat topics-to-move.json
{
  "topics": [{"topic": "lab42"}],
  "version":1
}
```

Generate a new assigment :

```
/opt/kafka/bin/kafka-reassign-partitions.sh --zookeeper zk:2181 --topics-to-move-json-file topics-to-move.json --broker-list "1002,1003,1004" --generate
```

`--broker-list` is the list of Kafka broker IDs you are targeting for the re-assignment. See the list in Zookeeper : [http://localhost:4550/?path=brokers%2Fids](http://localhost:4550/?path=brokers%2Fids)

This will generate a new assigment. You will have to put it in a file. Say new-assignment.json and tweak it if needed. At this point you can define exactly who will be the leader of every partitions.

Then you can apply the new assignement :

```
/opt/kafka/bin/kafka-reassign-partitions.sh --zookeeper zk:2181 --reassignment-json-file new-assignment.json --execute
```

You will be able to follow the status of the reassigment process with the following command :

```
/opt/kafka/bin/kafka-reassign-partitions.sh --zookeeper zk:2181 --reassignment-json-file new-assignment.json --verify
```
