package main

import (
	"log"

	"github.com/streadway/amqp"
)


func onFailError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	onFailError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err:= conn.Channel()
	onFailError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	onFailError(err, "Failed to declare a queue")

	body:="Hello World!"
	
	err = ch.Publish(
		"",   //exchange
		q.Name, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
	onFailError(err, "Failed to publish a message")
	}
	
	log.Printf(" [x] Sent %s", body)

}