package handlers

import (
	"fmt"

	"github.com/ResulShamuhammedov/howabot/api"
	"github.com/ResulShamuhammedov/howabot/models"
	"github.com/shomali11/slacker"
	"github.com/sirupsen/logrus"
)

func HandleWeather(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
	location := request.Param("location")

	data, err := api.GetWeather(location)
	if err != nil {
		eMsg := "Couldn't get weather, location: " + location
		logrus.WithError(err).Error(eMsg)
		return
	}

	if data.Location == "" {
		response.Reply("Sorry, I can't find this location...")
		return
	}

	message := fmt.Sprintf(models.ResponseFormat,
		data.Location,
		data.Weather[0].Main, data.Weather[0].Description,
		data.Main.Temperature-273.15, data.Main.Temperature,
		data.Main.FeelsLike-273.15, data.Main.FeelsLike,
		data.Main.Pressure,
		data.Main.Humidity,
		data.Visibility,
		data.Wind.Speed,
	)

	response.Reply(message)
}
