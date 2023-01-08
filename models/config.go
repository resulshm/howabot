package models

type Configuration struct {
	ListenPort        string `json:"listen_port"`
	OpenWeatherApiKey string `json:"open_weather_api_key"`
	SlackBotToken     string `json:"slack_bot_token"`
	SlackAppToken     string `json:"slack_app_token"`
}
