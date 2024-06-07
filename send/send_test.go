package send

import (
	"testing"

	"github.com/streadway/amqp"
)


func TestSend(t *testing.T){
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	onFailError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err:= conn.Channel()
	onFailError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"test_queue", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	onFailError(err, "Failed to declare a queue")

	body:="Test Message!"
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

}