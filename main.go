package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ResulShamuhammedov/howabot/api"
	"github.com/ResulShamuhammedov/howabot/handlers"
	"github.com/ResulShamuhammedov/howabot/models"
	"github.com/ResulShamuhammedov/howabot/utils"
	"github.com/gorilla/mux"
	"github.com/shomali11/slacker"
	"github.com/sirupsen/logrus"
)

func main() {

	err := utils.ReadConfig("config.json")
	if err != nil {
		eMsg := "Error in reading configuration"
		logrus.WithError(err).Error(eMsg)
		return
	}

	err = startBot(&utils.Config)
	if err != nil {
		logrus.WithError(err).Error("Couldn't start slack bot")
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

	r.HandleFunc("/{city}", handlers.HandleWeather)

	logrus.Info("Listen on port " + config.ListenPort)

	err := http.ListenAndServe(config.ListenPort, r)
	if err != nil {
		logrus.WithError(err).Error("Couldn't listen and serve")
		return err
	}

	return err
}

func startBot(config *models.Configuration) error {
	bot := slacker.NewClient(config.SlackBotToken, config.SlackAppToken)
	definition := &slacker.CommandDefinition{
		Description: "Weather",
		Examples:    []string{"weather London"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			city := request.Param("city")
			fmt.Println(city)
			data, err := api.GetWeather(city)
			if err != nil {
				eMsg := "Couldn't get weather, city: " + city
				logrus.WithError(err).Error(eMsg)
				return
			}
			out, _ := json.Marshal(data)
			response.Reply(string(out))
		},
	}

	bot.Command("weather {city}", definition)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		return err
	}
	return nil
}
