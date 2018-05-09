package main

import (
	"fmt"
	"github.com/morfeush22/go-tx/crc-calc/message"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
)

func handler(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("data")
	logger := log.WithFields(log.Fields{"data": data})
	logger.Debug("Received data")

	if len(data) != 0 {
		msg := message.NewMessage(data)

		js, err := msg.Marshalize()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logger.Fatal("Can not marshalize data")
			return
		}

		logger.WithField("crc", "0x" + fmt.Sprintf("%x", msg.CRC)).Debug("CRC has been calculated")
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
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

	serverPort := "8080"
	http.HandleFunc("/crc", handler)
	log.WithFields(log.Fields{
		"serverPort": serverPort,
	}).Info("Starting server")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil))
}
