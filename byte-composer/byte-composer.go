package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"io/ioutil"
	"encoding/json"
)

type config struct {
	crcCalcHost string
	crcCalcPort string
}

func (c *config) handler(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("data")
	logger := log.WithFields(log.Fields{"data": data})
	logger.Debug("Received data")

	var dat map[string]interface{}

	if len(data) != 0 {
		resp, err := http.Get("http://" + c.crcCalcHost + ":" + c.crcCalcPort + "/crc?data=" + data)
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

func main() {
	if os.Getenv("DEBUG") != "" {
		log.SetLevel(log.DebugLevel)
	}

	crcCalcHost := os.Getenv("CRC_CALC_HOST")
	crcCalcPort := os.Getenv("CRC_CALC_PORT")

	log.WithFields(log.Fields{
		"crcCalcHost": crcCalcHost,
		"crcCalcPort": crcCalcPort,
	}).Info()

	if crcCalcHost == "" || crcCalcPort == "" {
		log.Error("CRC_CALC_HOST or CRC_CALC_PORT empty")
		os.Exit(1)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Info("Exiting")
		os.Exit(0)
	}()

	serverPort := "8080"
	config := config{crcCalcHost, crcCalcPort}
	http.HandleFunc("/compose", config.handler)
	log.WithFields(log.Fields{
		"serverPort": serverPort,
	}).Info("Starting server")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil))
}
