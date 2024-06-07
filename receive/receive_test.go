package receive

import (
	"testing"

	"github.com/streadway/amqp"
)



func TestReceive(t *testing.T){
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	onFailError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
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
	if err != nil {
	onFailError(err, "Failed to declare a queue")
	}

	body:="Test Case!"
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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
	onFailError(err, "Failed to register a consumer")
	}
	
	received := false
	for d := range msgs {
		if string(d.Body) == body {
			received = true
			break
		}
	}

	if !received {
		t.Fatalf("Failed to receive the message")
	}


}

func TestMain(m *testing.M){
	m.Run()

	conn, _:= amqp.Dial("amqp://guest:guest@localhost:5672/")
	defer conn.Close()
	ch, _:= conn.Channel()
	defer ch.Close()

	conn.Close()

	ch.QueueDelete("test_hello", false, false, false)
}