package main

import (
	"github.com/streadway/amqp"
	"log"
	"strconv"
)

func failOnError(err error, msg string){
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
	}
}



func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"rpc_queue",
		false,
		false,
		false,
		false,
		nil,
		)

	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,
		0,
		false,
		)
	failOnError(err, "Failed to change a prefetch")

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func(){
		for d := range msgs {
			n, err := strconv.Atoi(string(d.Body))
			failOnError(err, "Failed to convert body to int")

			//log.Printf("Getted %d", n)
			n++

			err = ch.Publish(
				"",
				d.ReplyTo,
				false,
				false,
				amqp.Publishing{
					ContentType:     "text/plain",
					CorrelationId:   d.CorrelationId,
					Body:            []byte(strconv.Itoa(n)),
				})
			failOnError(err, "Failed to publish a message")

			d.Ack(false)
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")

	<-forever
}
