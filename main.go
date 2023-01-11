package main

import (
	"context"

	"github.com/ResulShamuhammedov/howabot/handlers"
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
		Examples:    []string{"weather in London"},
		Handler:     handlers.HandleWeather,
	}

	bot.Command("weather in {location}", definition)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		return err
	}
	return nil
}
