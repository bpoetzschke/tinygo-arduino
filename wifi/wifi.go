package wifi

import (
	"machine"
	"time"
	"tinygo.org/x/drivers/net"
	"tinygo.org/x/drivers/wifinina"
)

type Wifi struct {
	spi     machine.SPI
	adapter *wifinina.Device
}

func (w *Wifi) Setup() error {
	w.spi = machine.NINA_SPI

	err := w.spi.Configure(machine.SPIConfig{
		Frequency: 8 * 1e6,
		SDO:       machine.NINA_SDO,
		SDI:       machine.NINA_SDI,
		SCK:       machine.NINA_SCK,
	})
	if err != nil {
		return err
	}

	w.adapter = wifinina.New(w.spi,
		machine.NINA_CS,
		machine.NINA_ACK,
		machine.NINA_GPIO0,
		machine.NINA_RESETN)
	w.adapter.Configure()

	return nil
}

func (w *Wifi) ConnectToAP(ssid, pass string) error {
	time.Sleep(time.Second)
	connected := false
	for i := 0; i < 3; i++ {
		println("Connecting to " + ssid)
		err := w.adapter.ConnectToAccessPoint(ssid, pass, 10*time.Second)
		if err != nil {
			if err.Error() == net.ErrWiFiConnectTimeout.Error() {
				println("Connection error while connecting to wifi. Will retry.")
				time.Sleep(time.Second)
				continue
			}
			return err
		}
		connected = true
	}

	if !connected {
		return net.ErrWiFiConnectTimeout
	}

	println("Connected.")

	time.Sleep(500 * time.Millisecond)
	ip, _, _, err := w.adapter.GetIP()
	if err != nil {
		return err
	}
	println(ip.String())

	return nil
}
