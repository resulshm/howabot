package main

import (
	"net/http"

	"github.com/ResulShamuhammedov/howabot/models"
	"github.com/ResulShamuhammedov/howabot/utils"
	"github.com/sirupsen/logrus"
)

var Config models.Configuration

func main() {

	Config, err := utils.ReadConfig("config.json")
	if err != nil {
		eMsg := "Error in reading configuration"
		logrus.WithError(err).Error(eMsg)
		return
	}

	err = setupServer(Config)
	if err != nil {
		logrus.WithError(err).Error("Couldn't start server")
		return
	}

}

func setupServer(config *models.Configuration) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello ...."))
	})

	logrus.Info("Listen on port " + config.ListenPort)

	err := http.ListenAndServe(config.ListenPort, mux)
	if err != nil {
		logrus.WithError(err).Error("Couldn't listen and serve")
		return err
	}

	return err
}
