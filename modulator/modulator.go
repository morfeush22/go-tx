package main

import (
	"fmt"
	"github.com/morfeush22/go-tx/modulator/qpsk"
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
		mod := qpsk.Modulator{}
		sig := mod.Modulate([]byte(data))

		js, err := sig.Marshal()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logger.Fatal("Can not marshal data")
			return
		}

		logger.
			WithField("signalInPhase", "0x"+fmt.Sprintf("%x", sig.InPhase)).
			WithField("signalQuadrature", "0x"+fmt.Sprintf("%x", sig.Quadrature)).
			Debug("Signal has been modulated")
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
	http.HandleFunc("/modulator", handler)
	log.WithFields(log.Fields{
		"serverPort": serverPort,
	}).Info("Starting server")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil))
}
