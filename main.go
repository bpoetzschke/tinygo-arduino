package main

import (
	"fmt"
	"github.com/bpoetzschke/tinygo-arduino/epd"
	"github.com/bpoetzschke/tinygo-arduino/openweather"
	"github.com/bpoetzschke/tinygo-arduino/wifi"
	"machine"
	"time"
)

const (
	lat  = 42.365250
	long = -71.105011
)

var (
	w                 wifi.Wifi
	openWeatherClient = openweather.NewClient(lat, long, openWeatherMapAPIKey)
	currentWeather    openweather.CurrentWeatherData
	display           = epd.Display{}
)

func main() {
	waitSerial()

	err := w.Setup()
	PrintErrIfNotNil(err)

	err = w.ConnectToAP(ssid, pass)
	PrintErrIfNotNil(err)

	err = display.Setup()
	PrintErrIfNotNil(err)

	for {
		err := openWeatherClient.GetCurrentWeather(&currentWeather)
		PrintErrIfNotNil(err)

		println(fmt.Sprintf("Curr weather: %#v", currentWeather))

		println("Clear display")
		display.Clear()
		//println("Split screen")
		//display.SplitScreen()
		println("Display weather icon")
		//display.DisplayCurrentWeatherIcon(meteoicon.GetMeteoIcon(currentWeather.Weather.Icon), &meteoicon.MeteoIcon48pt)
		//display.DisplayCurrentWeatherCondition("Clouds")
		println("display")
		err = display.Display()
		PrintErrIfNotNil(err)
		println("display done")
		//currentWeather = nil

		time.Sleep(30 * time.Second)
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
