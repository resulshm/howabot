package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ResulShamuhammedov/howabot/models"
	"github.com/ResulShamuhammedov/howabot/utils"
	"github.com/sirupsen/logrus"
)

var OpenWeatherURL = "http://api.openweathermap.org/data/2.5/weather?APPID=%s&q=%s"

func GetWeather(city string) (*models.WeatherData, error) {

	respData, err := http.Get(fmt.Sprintf(OpenWeatherURL, utils.Config.OpenWeatherApiKey, city))
	if err != nil {
		eMsg := "Couldn't get response from OpenWeatherMap, city: " + city
		logrus.WithError(err).Error(eMsg)
		return nil, nil
	}

	defer respData.Body.Close()
	d := &models.WeatherData{}
	err = json.NewDecoder(respData.Body).Decode(&d)
	if err != nil {
		eMsg := "Couldn't decode response from OpenWeather Api"
		logrus.WithError(err).Error(eMsg)
		return nil, err
	}
	return d, nil
}
