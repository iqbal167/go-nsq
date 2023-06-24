package main

import (
	"encoding/json"
	"log"

	"github.com/nsqio/go-nsq"
)

// type (
// 	PaymentCallback struct {
// 		ID               string        `json:"id"`
// 		User             User          `json:"user"`
// 		PaymentMethod    PaymentMethod `json:"paymentMethod"`
// 		Amount           int           `json:"amount"`
// 		ReferenceType    string        `json:"referenceType"`
// 		RerefenceId      string        `json:"paymentId"`
// 		PaymentReference string        `json:"paymentReference"`
// 		Status           string        `json:"status"`
// 		CustomeName      string        `json:"customerName"`
// 	}

// 	User struct {
// 		ID string `json:"id"`
// 	}

// 	PaymentMethod struct {
// 		Code string `json:"code"`
// 	}
// )

type Profile struct {
	Username string `json:"username"`
}

type Message struct {
	Data Profile `json:"data"`
}

func main() {
	config := nsq.NewConfig()

	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	topic := "education-ppdb-user-profile"

	msg := Message{
		Data: Profile{
			Username: "20200582210001",
		},
	}

	payload, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	err = producer.Publish(topic, payload)
	if err != nil {
		log.Fatal(err)
	}
}
