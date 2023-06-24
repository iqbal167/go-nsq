package main

import (
	"log"
	"sync"

	"github.com/nsqio/go-nsq"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()

	topic := "users"
	channel := "utilities"

	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Panic("Could not create consumer")
	}
	//consumer.MaxInFlight defaults to 1

	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println("NSQ message received:")
		log.Println(string(message.Body))
		return nil
	}))

	err = consumer.ConnectToNSQLookupd("127.0.0.1:4161")
	if err != nil {
		log.Panic("Could not connect")
	}

	err = consumer.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect")
	}
	log.Println("Awaiting messages from NSQ topic \"My NSQ Topic\"...")
	wg.Wait()
}
