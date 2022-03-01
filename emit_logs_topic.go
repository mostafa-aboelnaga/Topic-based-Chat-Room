/*
USAGE:
1) you can push to ONLY ONE TOPIC, you may modify this code to emit to multiple topics, loop on publish
2) THIS CAN'T HAVE * OR #

go run emit_logs_topic.go anonymous.info # this sends to anonymous.info
go run emit_logs_topic.go kernel.info  # this sends to kernel.info
*/
package main

import (
	"bufio"                                      // to implement a buffered I/O, used in our inputReader
	"chatroom_topic_basaed_implementation/utils" // to use the error checking function implemented in utils/error.go
	"fmt"                                        // to implement formatted I/O with functions like in C language
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

	topic := severityFrom(os.Args) // a command line arg like web.info (this is a topic)
	// this is the topic I'm interested in

	inputReader := bufio.NewReader(os.Stdin) // defining a reader, so we can accept some input!

	for {
		// guiding the user into typing or publishing a message
		// in his specified topic
		fmt.Println(fmt.Sprintf("Enter a message to be published on (%s)", topic))

		// now reading the message, that the user would like to publish in such topic
		messageToBePublished, _, err := inputReader.ReadLine()
		body := string(messageToBePublished) // stringfying this byte-formatted received messsage

		// now publishing it in the specified topic!
		err = ch.Publish(
			"logs_topic", // exchange
			topic,        // routing key (THIS CAN'T HAVE * OR #)
			false,        // mandatory
			false,        // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		utils.FailOnError(err, "Failed to publish a message")

		// logging everything we publish into our specified topic
		// as in this format: [topic] Sent body
		// for example [egypt.politics] Sent sample_text_74
		log.Printf(" [%s] sent %s", topic, body)
	}
}

func severityFrom(args []string) string {
	// EXAMPLE: go run emit_log_topic.go anonymous.info, you can push to ONLY ONE TOPIC
	// (THIS CAN'T HAVE * OR #)
	// the default queue is anonymous.info
	// you may try something else
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "anonymous.info"
	} else {
		s = os.Args[1]
	}
	return s
}
