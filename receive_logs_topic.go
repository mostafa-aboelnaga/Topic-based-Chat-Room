/*
USAGE
go run receive_logs_topic.go *.info # receives from anything having info in the end
go run receive_logs_topic.go anonymous.*
go run receive_logs_topic.go kernel.error
*/
package main

import (
	"chatroom_topic_basaed_implementation/utils" // to use the error checking function implemented in utils/error.go
	"log"                                        // to implement simple logging
	"os"                                         // to provide a platform-independent interface to operating system functionality

	"github.com/streadway/amqp"
)

func main() {
	// connection statements
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type is now topic, in order to implement such a chat room that is based on topics
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	utils.FailOnError(err, "Failed to declare an exchange")

	// declaring the queue we would like to bind later
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)

	utils.FailOnError(err, "Failed to declare a queue")

	// checking whether we have some arguments or not, and if not,
	// we would exit and output some error indicating that we should have some arguments
	// referring to the topics we would like to listen to
	if len(os.Args) < 2 {
		log.Printf("Usage: %s [binding_key]...", os.Args[0])
		os.Exit(0)
	}

	// binding each argument as we go
	for _, binding := range os.Args[1:] {
		log.Printf("Binding queue %s to exchange %s with binding key %s",
			q.Name, "logs_topic", binding)
		err = ch.QueueBind(
			q.Name,       // queue name
			binding,      // binding key
			"logs_topic", // exchange
			false,
			nil)
		utils.FailOnError(err, "Failed to bind a queue")
	}

	// getting their copy of the published messages, consuming it
	publishedMessages, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	// to keep on the goroutine alive forever!
	// long live the blocking channel!
	forever := make(chan bool)

	go func() {
		for message := range publishedMessages {
			// we simply log each message along with who sent it, etc
			// we expect messages from the topics we binded to
			// we log it as in this format: [topic] Sent body
			// for example [egypt.politics] Sent sample_text_74

			log.Printf(" [%s] sent %s", message.RoutingKey, message.Body)
			// log.Printf(" [x] %s", d.Body)
			// log.Printf(string(d.Body))
		}
	}()

	log.Printf(" //ATTENTION// Waiting for logs. To exit press CTRL+C!")
	<-forever
}
