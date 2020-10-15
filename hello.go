package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/streadway/amqp"
)


	type Payload struct {
		TimeStamp int64 `json:"time_stamp"`
	}
func main () {

	var p Payload
	for{
		p.TimeStamp = time.Now().Unix()
		time.Sleep(time.Second ^1)
	}

}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Connection() {
	conn, err := amqp.Dial("amqp://Maxine:daddy@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	FailOnError(err, "Failed to declare a queue")

	body := json
	err = ch.Publish(
		"",
		q.Name{
			false,
			false,
			ampq.Publishing{
				ContentType: "json",
				Body: []byte(body),
			}
		)
	FailOnError("Failed to publish a message")
}


