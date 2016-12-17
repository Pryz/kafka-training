package main

import(
	"fmt"
	"flag"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	broker string
	topic string
	value string
)

func main() {

	flag.StringVar(&broker, "broker", "localhost", "Kafka broker hostname")
	flag.StringVar(&topic, "topic", "", "Kafka topic name")
	flag.StringVar(&value, "value", "", "Write this data to Kafka")
	flag.Parse()

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		fmt.Printf("Failed to create producer : %s\n", err)
	}

	fmt.Printf("Create producer %v\n", p)

	// .Events channel is used.
	deliveryChan := make(chan kafka.Event)

	value := "Hello Go!"
	err = p.Produce(&kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Value: []byte(value)}, deliveryChan)

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)
}
