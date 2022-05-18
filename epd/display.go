package epd

import (
	"github.com/bpoetzschke/tinygo-arduino/freesans"
	"image/color"
	"machine"
	"tinygo.org/x/drivers/waveshare-epd/epd2in13x"
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
}

func (d *Display) Setup() error {

	err := machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8000000,
	})
	if err != nil {
		return err
	}

	d.epDisplay = epd2in13x.New(machine.SPI0, machine.D10, machine.D9, machine.D8, machine.D7)
	d.epDisplay.Configure(epd2in13x.Config{})
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

func (d *Display) DisplayCurrentWeatherIcon(icon string, font tinyfont.Fonter) {
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
	println("Displaying icon:", icon)
	// the icon needs to be positioned centered in the bounding box of the font ideally
	tinyfont.WriteLineRotated(&d.epDisplay, font, displayHeight-2, 1, icon, colorBlack, tinyfont.ROTATION_90)
}

func (d *Display) DisplayCurrentWeatherCondition(cond string) {
	tinyfont.WriteLineRotated(&d.epDisplay, &freesans.Regular12pt7b, (displayHeight/2)-5+17, 1, cond, colorBlack, tinyfont.ROTATION_90)
}

func (d *Display) Display() error {
	err := d.epDisplay.Display()
	if err != nil {
		return err
	}

	d.epDisplay.WaitUntilIdle()
	return nil
}
