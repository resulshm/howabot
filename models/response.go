package models

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type WeatherData struct {
	City string `json:"name"`
	Main struct {
		Temperature float64 `json:"temp"`
	} `json:"main"`
}
