# Kafka training

The purpose of this training is to give you a starting point to use and experiment with Kafka.

## Resources

Some interesting resources :

* The official Kafka documentation : [https://kafka.apache.org/documentation/](https://kafka.apache.org/documentation/). You should at least read the [getting started](http://kafka.apache.org/documentation/#gettingStarted<Paste>) and [design](http://kafka.apache.org/documentation/#design) sections. 
* A training Deck of Kafka 0.8 : [http://www.slideshare.net/miguno/apache-kafka-08-basic-training-verisign](http://www.slideshare.net/miguno/apache-kafka-08-basic-training-verisign)
* Monitoring Kafka from Datadog folks : [https://www.datadoghq.com/blog/monitoring-kafka-performance-metrics/](https://www.datadoghq.com/blog/monitoring-kafka-performance-metrics/)
* Interesting reading for a Production deployment from Confluent : [http://docs.confluent.io/3.1.1/schema-registry/docs/deployment.html?highlight=production](http://docs.confluent.io/3.1.1/schema-registry/docs/deployment.html?highlight=production)
* Kafka replication - a lesson in operational simplicity : [https://www.confluent.io/blog/hands-free-kafka-replication-a-lesson-in-operational-simplicity](https://www.confluent.io/blog/hands-free-kafka-replication-a-lesson-in-operational-simplicity)

You don't have to read ALL of these before starting playing with the lab. But these are definitelly articles where you will find a lot of information to answer your question and learn more about Kafka.

### Requirements 

* Install Docker and Docker Compose
* Clone this repository on your laptop

### Start cluster

```
cd docker
```

You will first need to change the docker-compose.yml file and change KAFKA_ADVERTISED_HOST_NAME with the IP of your Docker Host.
If you are using `Docker for Mac` it will probably be your local IP, if you are using Boot2docker, it will be the IP for the VirtualBox instance, etc

Start containers :

```
cd docker
docker-compose up
```

This will start :

* A Zookeeper node listening on port 2181
* A Kafka broker
* A Zookeeper web app to browse the ZK content listening on 4550. ([http://localhost:4550](http://localhost:4550)) from your web browser.

The Input connection string for the ZK browser will be : zk:2181.

### "Connect" to Kafka

Since we are using Docker containers for this training, we will execute a bash process inside the Kafka container to perform all the actions.

```
docker exec -it docker_kafka_1 bash

ps auxwww | grep kafka
```

You should see a Kafka process up and running.

### Create lab42 topic

Let's create a simple topic :

```
/opt/kafka/bin/kafka-topics.sh --zookeeper zk:2181 --create --replication-factor 1 --partitions 1 --topic lab42
```

### Produce to lab42 topic

From here, you should already be able to write something in the new topic :

```
echo "helloworld" | \
/opt/kafka/bin/kafka-console-producer.sh --broker-list localhost:9092 --topic lab42
```

### Consume from lab42 topic

Easy to produce but also easy to read from that topic :

```
/opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic lab42 --from-beginning
```

### Audit cluster (list topics, leaders, ...)

`--describe` command will list you all the current topic and the leader and replicas of every partitions :

```
/opt/kafka/bin/kafka-topics.sh --zookeeper zk:2181 --describe
```

In production you will want to use the option `--under-replicated` to list all the under replicated partitions.

Really good article about replication : [https://www.confluent.io/blog/hands-free-kafka-replication-a-lesson-in-operational-simplicity](https://www.confluent.io/blog/hands-free-kafka-replication-a-lesson-in-operational-simplicity)

From here, you can start playing with the `consumer` and `production` Go app example here : [https://github.com/Pryz/kafka-training/tree/master/goclient](https://github.com/Pryz/kafka-training/tree/master/goclient)

### Add two brokers in the cluster

Let's scale the cluster by adding 2 more Kafka brokers :

```
docker-compose scale kafka=3
```

Within Zookeeper (using the ZK browser), you will get details about each Broker : [http://localhost:4550/?path=brokers%2Fids](http://localhost:4550/?path=brokers%2Fids)

### Add partitions to lab42 topic

We have 3 brokers now, let's scale the topic and add more partitions :

```
/opt/kafka/bin/kafka-topics.sh --zookeeper zk:2181 --alter --topic lab42 --partitions 3
```

### Look at the current topic repartition

From here, every broker should be leader of one partitions :

```
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

### Increase the replication factor of a topic 

Usually you don't do this task in Production. You directly define the replication factor when you create  the topic.

But if you have to increase the replication factor, you will need to use the reassigment tool. 

```
$ cat replication-factor.json
{"version":1,
"partitions":[
  {"topic":"lab42","partition":0,"replicas":[1001,1002,1003]},
  {"topic":"lab42","partition":1,"replicas":[1002,1003,1001]},
  {"topic":"lab42","partition":2,"replicas":[1003,1001,1002]}
]
}
```

Replicas is a list of broker IDs. As said before, you will find all the Broker IDs in ZK : [http://localhost:4550/?path=brokers%2Fids](http://localhost:4550/?path=brokers%2Fids)

Here, we will increase the replication factor to 3 and properly balance partitions across the cluster.

```
/opt/kafka/bin/kafka-reassign-partitions.sh --zookeeper zk:2181 --reassignment-json-file replication-factor.json --execute
```

To make sure it's done :

```
/opt/kafka/bin/kafka-reassign-partitions.sh --zookeeper zk:2181 --reassignment-json-file replication-factor.json --verify
```
