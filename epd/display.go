package epd

import (
	"fmt"
	"github.com/bpoetzschke/tinygo-arduino/epd2in13x"
	"github.com/bpoetzschke/tinygo-arduino/freesans"
	"github.com/bpoetzschke/tinygo-arduino/openweather"
	"image/color"
	"machine"
	"strings"
	"time"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
)

const (
	displayWidth  = 212
	displayHeight = 104
)

var (
	colorBlack = color.RGBA{1, 1, 1, 255}
	colorRed   = color.RGBA{1, 0, 0, 255}
	colorWhite = color.RGBA{1, 0, 0, 255}
)

type Display struct {
	epDisplay  epd2in13x.Device
	colorBlack color.RGBA
	debug      bool
}

func (d *Display) Setup(debug bool) error {

	err := machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8000000,
	})
	if err != nil {
		return err
	}

	d.epDisplay = epd2in13x.New(machine.SPI0, machine.D10, machine.D9, machine.D8, machine.D7)
	d.epDisplay.Configure(epd2in13x.Config{})
	d.debug = debug
	return nil
}

func (d *Display) Clear() {
	d.epDisplay.ClearBuffer()
	d.epDisplay.ClearDisplay()
}

func (d *Display) SplitScreen() {
	tinydraw.Line(&d.epDisplay, 0, displayWidth/2, displayHeight, displayWidth/2, colorBlack)
	tinydraw.Line(&d.epDisplay, displayHeight/2, 0, displayHeight/2, displayWidth, colorBlack)
}

func (d *Display) DisplayCurrentWeather(currWeather *openweather.CurrentWeatherData, meteoIcon string, font *tinyfont.Font) {
	// long side of display is y and short one is x
	/*
			--> x
		|	+-------+
		V	|       |
		y	|       |
			|       |
			|       |
			|       |
			+-------+
	*/
	glyph := font.GetGlyph(rune(meteoIcon[0]))
	maxWidth := font.BBox[0]
	maxHeight := font.BBox[1]

	weatherIconXOrigin := int16(displayHeight - 25)
	weatherIconYOrigin := int16(1)

	glyphXOffset := int16(maxHeight) - int16(glyph.Info().Height)
	if glyphXOffset > 0 {
		glyphXOffset /= 2
	}

	glyphYOffset := int16(maxWidth) - int16(glyph.Info().Width)
	if glyphYOffset > 0 {
		glyphYOffset /= 2
	}

	xPos := weatherIconXOrigin - glyphXOffset
	yPos := weatherIconYOrigin + glyphYOffset

	condFont := &freesans.Regular9pt7b
	//condFont := &freesans.FreeSans18pt
	//tempFont := &freesans.Regular9pt7b

	currTime := time.Unix(int64(currWeather.DT)-14400, 0)
	currDayStr := currTime.Format("Mon, 02 Jan 2006")
	tinyfont.WriteLineRotated(&d.epDisplay, condFont, displayHeight-1-int16(condFont.BBox[1]), 1, currDayStr, colorBlack, tinyfont.ROTATION_90)

	if d.debug {
		tinydraw.Rectangle(&d.epDisplay, weatherIconXOrigin-int16(maxHeight), weatherIconYOrigin, int16(maxWidth), int16(maxHeight), colorBlack)
	}

	tinyfont.WriteLineRotated(&d.epDisplay, font, xPos, yPos, meteoIcon, colorBlack, tinyfont.ROTATION_90)

	condXOrigin := weatherIconXOrigin - int16(maxHeight) - int16(condFont.BBox[1]) - 5
	tinyfont.WriteLineRotated(&d.epDisplay, condFont, condXOrigin, 1, strings.Title(currWeather.Weather.Description), colorBlack, tinyfont.ROTATION_90)
	d.showPageIndicator(4, 1)

	tinyfont.WriteLineRotated(&d.epDisplay, condFont, 65, 52, fmt.Sprintf("Temperature: %d ºC", currWeather.Main.Temp), colorBlack, tinyfont.ROTATION_90)
	tinyfont.WriteLineRotated(&d.epDisplay, condFont, 48, 52, "fo%o", colorBlack, tinyfont.ROTATION_90)
}

func (d *Display) showPageIndicator(totalPages, currentPage int8) {
	if currentPage > totalPages {
		currentPage = totalPages
	} else if currentPage < 1 {
		currentPage = 1
	}

	circleRadius := int16(4)
	circleSpacing := (2 * circleRadius) + 2
	circleXPos := displayHeight - 7 - circleRadius
	circleYStart := displayWidth - 4 - (int16(totalPages-1) * circleSpacing) - 2*circleRadius

	for i := int8(0); i < totalPages; i++ {
		if (currentPage - 1) == i {
			tinydraw.FilledCircle(&d.epDisplay, circleXPos, circleYStart+(int16(i)*circleSpacing), circleRadius, colorRed)
		} else {
			tinydraw.Circle(&d.epDisplay, circleXPos, circleYStart+(int16(i)*circleSpacing), circleRadius, colorBlack)
		}
	}
}

func (d *Display) Display() error {
	err := d.epDisplay.Display()
	if err != nil {
		return err
	}

	d.epDisplay.WaitUntilIdle()
	return nil
}

func (d *Display) DisplayFrame() {
	d.epDisplay.DisplayFrame()

	d.epDisplay.WaitUntilIdle()
}

func (d *Display) SetRectColor() error {
	width := int16((212 / 2) / 8)
	height := int16(104 / 2)
	buf := make([]byte, width*height)

	for i := 0; i < len(buf); i++ {
		buf[i] = 0x00
	}

	return d.epDisplay.SetDisplayRectColor(buf, 0, 0, width*8, height, epd2in13x.BLACK)
}
