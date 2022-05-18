package openweather

import (
	"fmt"
	"github.com/buger/jsonparser"
	"net/url"
	"strconv"
	"strings"
	"time"
	"tinygo.org/x/drivers/net"
	"tinygo.org/x/drivers/net/http"
	"tinygo.org/x/drivers/net/tls"
)

const (
	server                 = "api.openweathermap.org"
	protocol               = "https"
	baseAPIUrl             = "/data/2.5/"
	currentWeatherEndpoint = "weather"
)

var (
	buf [1024]byte
)

type Client interface {
	GetCurrentWeather(data *CurrentWeatherData) error
}

func NewClient(lat, long float64, apiKey string) Client {
	c := client{
		lat:    lat,
		long:   long,
		apiKey: apiKey,
	}

	return &c
}

type client struct {
	lat        float64
	long       float64
	apiKey     string
	connection net.Conn
}

func (c *client) GetCurrentWeather(currWeather *CurrentWeatherData) error {
	currentWeatherUrl := fmt.Sprintf("%s://%s%s%s?lat=%f&lon=%f&appid=%s&units=metric", protocol, server, baseAPIUrl, currentWeatherEndpoint, c.lat, c.long, c.apiKey)
	println(currentWeatherUrl)

	start, end, err := c.makeHTTPSRequest(currentWeatherUrl)

	if err != nil {
		return err
	}

	firstWeatherData, _, _, err := jsonparser.Get(buf[start:end], "weather", "[0]")
	if err != nil {
		return err
	}
	err = c.parseWeather(firstWeatherData, &currWeather.Weather)
	if err != nil {
		return err
	}

	mainData, _, _, err := jsonparser.Get(buf[start:end], "main")
	if err != nil {
		return err
	}
	err = c.parseMain(mainData, &currWeather.Main)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) makeHTTPSRequest(rawUrl string) (int, int, error) {
	println("HTTPS request to", rawUrl)
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return -1, -1, err
	}

	conn, err := tls.Dial("tcp", parsedUrl.Host, nil)
	retry := 0
	for ; err != nil; conn, err = tls.Dial("tcp", parsedUrl.Host, nil) {
		retry++
		if retry > 10 {
			return -1, -1, fmt.Errorf("Connection failed: %s", err.Error())
		}
		time.Sleep(1 * time.Second)
	}

	path := parsedUrl.Path
	if path == "" {
		path = "/"
	}

	if parsedUrl.RawQuery != "" {
		path += "?" + parsedUrl.RawQuery
	}

	fmt.Fprintln(conn, "GET "+path+" HTTP/1.1")
	fmt.Fprintln(conn, "Host:", parsedUrl.Host)
	fmt.Fprintln(conn, "User-Agent: TinyGo")
	fmt.Fprintln(conn, "Connection: close")
	fmt.Fprintln(conn)

	time.Sleep(time.Second)
	println("Reading conn")

	defer conn.Close()
	var offset = 0
	// TODO consider reading the payload at hoc instead of writing it to a buffer
	for n, err := conn.Read(buf[offset:]); n > 0; n, err = conn.Read(buf[offset:]) {
		if err != nil {
			return -1, -1, err
		}
		offset += n
	}

	var start, end int
	var statusStart, statusEnd int
	for i := 0; i < len(buf); i++ {
		if (buf[i] == 0x0a || buf[i] == 0x0d) /*\n || \r*/ && statusEnd == 0 {
			statusEnd = i
		}
		if buf[i] == 0x7b && start == 0 /*{*/ {
			start = i
		}
		if buf[i] == 0x7d /*}*/ {
			end = i
		}
	}

	if statusEnd == 0 {
		return -1, -1, fmt.Errorf("unable to read http status row")
	}

	statusParts := strings.Split(string(buf[statusStart:statusEnd]), " ")
	if len(statusParts) != 3 {
		return -1, -1, fmt.Errorf("Unknown http status row")
	}
	statusInt, err := strconv.Atoi(statusParts[1])
	if err != nil {
		return -1, -1, err
	}
	if statusInt != http.StatusOK {
		return -1, -1, fmt.Errorf("unexpected http status %d", statusInt)
	}

	println("Reading done")
	return start, end, nil
}

func (c *client) parseWeather(data []byte, target *Weather) error {
	var err error
	target.ID, err = jsonparser.GetInt(data, "id")
	if err != nil {
		return err
	}
	target.Main, err = jsonparser.GetString(data, "main")
	if err != nil {
		return err
	}
	target.Description, err = jsonparser.GetString(data, "description")
	if err != nil {
		return err
	}
	target.Icon, err = jsonparser.GetString(data, "icon")
	if err != nil {
		return err
	}

	return nil
}

func (c *client) parseMain(data []byte, target *MainWeatherData) error {
	var err error
	target.Temp, err = jsonparser.GetFloat(data, "temp")
	if err != nil {
		return err
	}
	target.FeelsLike, err = jsonparser.GetFloat(data, "feels_like")
	if err != nil {
		return err
	}
	target.TempMin, err = jsonparser.GetFloat(data, "temp_min")
	if err != nil {
		return err
	}
	target.TempMax, err = jsonparser.GetFloat(data, "temp_max")
	if err != nil {
		return err
	}
	target.Pressure, err = jsonparser.GetFloat(data, "pressure")
	if err != nil {
		return err
	}
	target.Humidity, err = jsonparser.GetFloat(data, "humidity")
	if err != nil {
		return err
	}
	return nil
}
