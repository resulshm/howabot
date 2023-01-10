package models

const ResponseFormat = `	
	City: %s
	Weather: %s (%s)
	Temperature: %0.2f°C or %0.2fK
	Feels like: %0.2f°C or %0.2fK
	Pressure: %0.f hPa
	Humidity: %0.f %%
	Visibility: %0.f m
	Wind speed: %0.2f meter/sec`

type WeatherData struct {
	City    string `json:"name"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temperature float64 `json:"temp"`
		FeelsLike   float64 `json:"feels_like"`
		Pressure    float64 `json:"pressure"`
		Humidity    float64 `json:"humidity"`
	} `json:"main"`
	Visibility float64 `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}
