package main

import (
	"log"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	brokerAddress := "localhost:9092"
	topic := "topic1"
	numMessages := 10000
	message := "Hello"

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Key = sarama.StringEncoder("Key")
	msg.Value = sarama.StringEncoder(message)

	producer, err := sarama.NewSyncProducer([]string{brokerAddress}, config)

	if err != nil {
		log.Printf("producer close, err: %s", err)
		return
	}

	defer producer.Close()

	startTime := time.Now()

	for i := 1; i <= numMessages; i++ {
		pid, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Printf("send failed, %s", err)
			panic(err)
		}
		log.Printf("pid: %d, offset: %d, topic: %s, key: %s, msg: %s\n", pid, offset, msg.Topic, msg.Key, msg.Value)
	}

	log.Printf("sent %d messages to Kafka topic '%s'\n", numMessages, topic)

	diff := time.Now().Sub(startTime)
	log.Printf("time taken: %s", diff)
}
