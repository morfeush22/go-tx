package main

import (
	"fmt"
	"github.com/morfeush22/go-tx/modulator/qpsk"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"os"
	"os/signal"
	"time"
)

type config struct {
	queueHost string
	queuePort string
}

func handleMessages(consumer <-chan amqp.Delivery, channel *amqp.Channel) error {
	mod := qpsk.Modulator{}

	for m := range consumer {
		sig := mod.Modulate([]byte(m.Body))
		log.WithField("signalInPhase", "0x"+fmt.Sprintf("%x", sig.InPhase)).
			WithField("signalQuadrature", "0x"+fmt.Sprintf("%x", sig.Quadrature)).
			Debug("Signal has been modulated")

		js, err := sig.Marshal()
		if err != nil {
			log.Error("Can not marshal data")
			continue
		}

		err = channel.Publish(
			"",
			m.ReplyTo,
			false,
			false,
			amqp.Publishing{
				ContentType:   "application/json",
				CorrelationId: "",
				Body:          js,
			})
		if err != nil {
			log.Error("Can not send data to queue")
			continue
		}
	}

	return nil
}

func (c *config) listen() error {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://guest:guest@%s:%s/", c.queueHost, c.queuePort))
	if err != nil {
		log.WithField("queueHost", c.queueHost).
			WithField("queuePort", c.queuePort).
			Error("Can not connect to AMQP server")
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
		"modulator",
		false,
		false,
		false,
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

	return handleMessages(consumer, channel)
}

func main() {
	if os.Getenv("DEBUG") != "" {
		log.SetLevel(log.DebugLevel)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Info("Exiting")
		os.Exit(0)
	}()

	config := config{os.Getenv("QUEUE_HOST"), os.Getenv("QUEUE_PORT")}
	for err := config.listen(); err != nil; err = config.listen() {
		time.Sleep(1 * time.Second)
	}
}
