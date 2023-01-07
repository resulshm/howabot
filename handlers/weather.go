package handlers

import (
	"errors"
	"net/http"

	"github.com/ResulShamuhammedov/howabot/api"
	"github.com/ResulShamuhammedov/howabot/utils"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func HandleWeather(w http.ResponseWriter, r *http.Request) {
	city := mux.Vars(r)["city"]
	if len(city) == 0 {
		eMsg := "city name must be given"
		err := errors.New(eMsg)
		logrus.WithError(err)
		utils.SendResponse(w, 400, eMsg)
		return
	}

	data, err := api.GetWeather(city)
	if err != nil {
		eMsg := "error in api.GetWeather, city: " + city
		err := errors.New(eMsg)
		logrus.WithError(err)
		utils.SendResponse(w, 500, eMsg)
		return
	}

	utils.SendResponse(w, 200, data)
	logrus.Info("Successfully hadled request. City: " + city)
}
