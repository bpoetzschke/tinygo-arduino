package openweather

type CurrentWeatherData struct {
	Weather Weather
	Main    MainWeatherData
}

type MainWeatherData struct {
	Temp      float64
	FeelsLike float64
	TempMin   float64
	TempMax   float64
	Pressure  float64
	Humidity  float64
}

type Weather struct {
	ID          int64
	Main        string
	Description string
	Icon        string
}
