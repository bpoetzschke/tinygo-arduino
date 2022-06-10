package main

import (
	"github.com/bpoetzschke/tinygo-arduino/epd"
	"github.com/bpoetzschke/tinygo-arduino/openweather"
	"github.com/bpoetzschke/tinygo-arduino/wifi"
	"machine"
	"time"
)

const (
	lat   = 42.365250
	long  = -71.105011
	DEBUG = false
)

var (
	w                 wifi.Wifi
	openWeatherClient = openweather.Client{Lat: lat, Long: long, ApiKey: openWeatherMapAPIKey}
	currentWeather    openweather.CurrentWeatherData
	display           = epd.Display{}
)

func main() {
	waitSerial()

	openWeatherClient.Setup()
	err := w.Setup()
	PrintErrIfNotNil(err)

	err = w.ConnectToAP(ssid, pass)
	PrintErrIfNotNil(err)

	err = display.Setup(DEBUG)
	PrintErrIfNotNil(err)

	for {
		//err := openWeatherClient.GetCurrentWeather(&currentWeather)
		//PrintErrIfNotNil(err)

		display.Clear()
		//err := display.SetRectColor()
		//PrintErrIfNotNil(err)
		//display.DisplayCurrentWeather(&currentWeather, meteoicon.GetMeteoIcon(currentWeather.Weather.Icon), &meteoicon.MeteoIcon48pt)
		//display.DisplayFrame()
		err := display.Display()
		PrintErrIfNotNil(err)
		println("display done")

		time.Sleep(time.Minute)
	}
}

// Wait for user to open serial console
func waitSerial() {
	for !machine.Serial.DTR() {
		time.Sleep(100 * time.Millisecond)
	}
}

func PrintErrIfNotNil(err error) {
	if err == nil {
		return
	}

	for {
		println(err)
		time.Sleep(time.Second)
	}
}
