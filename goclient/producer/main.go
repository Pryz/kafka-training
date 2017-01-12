package main

import(
	"log"
	"flag"

	"github.com/Shopify/sarama"
)

var (
	brokerList []string
	topic string
	value string
)

func main() {

	flag.StringVar(&topic, "topic", "", "Kafka topic name")
	flag.StringVar(&value, "value", "", "Write this data to Kafka")
	flag.Parse()

	brokerList = flag.Args()
	log.Printf("brokerList : %v\n", brokerList)

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(value),
	})
	if err != nil {
		log.Fatalln("Failed to produce message:", err)
	}
	log.Printf("Topic: %s, Partition: %d, Offset: %d, Message: %s\n", topic, partition, offset, value)

}
