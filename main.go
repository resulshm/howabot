package main

import (
	"net/http"

	"github.com/ResulShamuhammedov/howabot/handlers"
	"github.com/ResulShamuhammedov/howabot/models"
	"github.com/ResulShamuhammedov/howabot/utils"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {

	err := utils.ReadConfig("config.json")
	if err != nil {
		eMsg := "Error in reading configuration"
		logrus.WithError(err).Error(eMsg)
		return
	}

	err = setupServer(&utils.Config)
	if err != nil {
		logrus.WithError(err).Error("Couldn't start server")
		return
	}

}

func setupServer(config *models.Configuration) error {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello ..."))
	})
	r.HandleFunc("/{city}", handlers.HandleWeather)

	logrus.Info("Listen on port " + config.ListenPort)

	err := http.ListenAndServe(config.ListenPort, r)
	if err != nil {
		logrus.WithError(err).Error("Couldn't listen and serve")
		return err
	}

	return err
}
