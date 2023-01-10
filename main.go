package main

import (
	"context"
	"fmt"

	"github.com/ResulShamuhammedov/howabot/api"
	"github.com/ResulShamuhammedov/howabot/models"
	"github.com/ResulShamuhammedov/howabot/utils"
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

			message := fmt.Sprintf(models.ResponseFormat,
				data.City,
				data.Weather[0].Main, data.Weather[0].Description,
				data.Main.Temperature-273.15, data.Main.Temperature,
				data.Main.FeelsLike-273.15, data.Main.FeelsLike,
				data.Main.Pressure,
				data.Main.Humidity,
				data.Visibility,
				data.Wind.Speed,
			)
			// out, _ := json.Marshal(data)
			response.Reply(message)
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
