package main

import (
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func failOnError(err error, msg string){
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
	}
}

func genRandString(l int) string {
	bytes := make([]byte, l)

	for i:=0; i < l; i++ {
		bytes[i] = byte(randInt(65,90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func sender(n int) (res int, err error){
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
		)
	failOnError(err, "Failed to declare queue")

	corrId := genRandString(32)

	err = ch.Publish(
		"",
		"rpc_queue",
		false,
		false,
		amqp.Publishing{
			ContentType:     "text/plain",
			CorrelationId:   corrId,
			ReplyTo:         q.Name,
			Body:            []byte(strconv.Itoa(n)),
		})

	failOnError(err, "Failed to publish a message")

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
		)
	failOnError(err, "Failed to register a consumer")

	go func(){
		for d := range msgs {
			if d.CorrelationId == corrId {
				res, err = strconv.Atoi(string(d.Body))
				failOnError(err, "Failed to convert body to string")
				break
			}
		}
	}()

	return
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	//n := bodyFrom(os.Args)

	for i:=0; i < 100; i++ {
		//log.Printf(" [x] Requesting %d + 1", i)
		t := time.Now()
		_, err := sender(i)
		d := time.Since(t)
		failOnError(err, "Failed to handle RPC request")
		log.Printf("Spended time %v milliseconds", d.Milliseconds())
		//log.Printf(" [.] Got %d", res)
	}
//asdf

}

func bodyFrom(args []string) int{
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "30"
	} else {
		s = strings.Join(args[1:], " ")
	}
	n, err := strconv.Atoi(s)
	failOnError(err, "Failed to convert arg to integer")
	return n
}
