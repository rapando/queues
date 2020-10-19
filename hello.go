package main

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	_ "net/http"
	"os"
	"time"
)


	type Payload struct {
		TimeStamp int64 `json:"time_stamp"`
	}
func main () {
	log.Println("Starting...")
	var counter = 0
	_= godotenv.Load()

	channel:= connecttoQueue()
	defer channel.Close

	var p Payload
	log.Println("Looks good")

	for{
		p.TimeStamp = time.Now().Unix()
		log.Printf("Step %d", counter)
		go publishToQueue (p, channel)
		time.Sleep(time.Second ^1)
		counter ++
	}

}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func ConnectToQueue() *ampq.Channel {
	log.Println("Connecting to queue")
	conn, err := amqp.Dial(os.Getenv("Q_URL"))
	FailOnError(err, "Failed to connect to RabbitMQ")

	log.Println("Creating a channel")
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	return ch
}

func PublishToQueue(p Payload, ch *amqp.Channel){
	queueName:os.Getenv("Q_NAME")
	jsonByte, _:= json.Marshal(p)
	log.Println(">>> Publishing", string(jsonByte), "to queue", queueName)

	err:= ch.Publish ("", queueName, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body: jsonByte,
	})
	FailOnError(err, "Unable to publish to queue")
	log.Println("<<< Done")
}

func ReceiveFromQueue(){
	log.Println("Connecting to queue")
	conn, err := amqp.Dial(os.Getenv("Q_URL"))
	FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")

	forever:= make(chan bool)
	go func() {
		for d:= range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

