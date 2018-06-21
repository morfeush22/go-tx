package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type config struct {
	queueHost string
	queuePort string
}

func (c *config) handler(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("data")
	logger := log.WithFields(log.Fields{"data": data})
	logger.Debug("Received data")

	var dat map[string]interface{}

	if len(data) != 0 {
		resp, err := http.Get("http://" + c.queueHost + ":" + c.queuePort + "/crc?data=" + data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logger.Error("Can not get CRC for data")
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err := json.Unmarshal(body, &dat); err != nil {
			panic(err)
		}

		logger.WithFields(log.Fields{
			"body": dat["data"],
		}).Info("Received data")
	}
}

func writeMessage(channel *amqp.Channel, replyQueue string, message string) error {
	err := channel.Publish(
		"",
		"crc-calc",
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: "",
			ReplyTo:       replyQueue,
			Body:          []byte(message),
		})
	if err != nil {
		log.Error("Failed to publish a message")
		return err
	}

	return nil
}

func consumeResponse(consumer <-chan amqp.Delivery) {
	select {
	case r := <-consumer:
		var data map[string]interface{}
		if err := json.Unmarshal(r.Body, &data); err != nil {
			log.Error("Can not unmarshal data")
			return
		}
		fmt.Println(data)

	case <-time.After(1 * time.Second):
		log.Error("Consumer timed out, waited too long for response")
	}
}

func (c *config) sendMessage(message string) error {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://guest:guest@%s:%s/", c.queueHost, c.queuePort))
	if err != nil {
		log.Error("Can not connect to AMQP server")
		return err
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Error("Can not open the AMQP channel")
		return err
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		log.Error("Can not declare AMQP queue")
		return err
	}

	consumer, err := channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Error("Can not register AMQP consumer")
		return err
	}

	err = writeMessage(channel, queue.Name, message)
	if err == nil {
		consumeResponse(consumer)
	}

	return nil
}

func main() {
	if os.Getenv("DEBUG") != "" {
		log.SetLevel(log.DebugLevel)
	}

	queueHost := os.Getenv("QUEUE_HOST")
	queuePort := os.Getenv("QUEUE_PORT")

	log.WithFields(log.Fields{
		"queueHost": queueHost,
		"queuePort": queuePort,
	}).Debug()

	if queueHost == "" || queuePort == "" {
		log.Error("QUEUE_HOST or QUEUE_PORT empty")
		os.Exit(1)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Info("Exiting")
		os.Exit(0)
	}()

	if len(os.Args) != 2 {
		log.Error("Specify message to send as first argument")
		os.Exit(2)
	}

	msg := os.Args[1]
	log.WithField("message", msg).Debug()

	config := config{queueHost, queuePort}
	config.sendMessage(msg)
}
