package openweather

type CurrentWeatherData struct {
	Weather Weather
	Main    MainWeatherData
	DT      int32
}

type MainWeatherData struct {
	Temp     int16
	TempMin  int16
	TempMax  int16
	Humidity int16
}

type Weather struct {
	ID          int16
	Main        string
	Description string
	Icon        string
}
