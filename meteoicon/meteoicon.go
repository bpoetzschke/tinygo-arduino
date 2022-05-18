package meteoicon

var lookupMap = map[string]string{
	"01d": "B",
	"01n": "C",
	"02d": "H",
	"02n": "4",
	"03d": "N",
	"03n": "5",
	"04d": "Y",
	"04n": "%",
	"09d": "R",
	"09n": "8",
	"10d": "Q",
	"10n": "7",
	"11d": "P",
	"11n": "6",
	"13d": "W",
	"13n": "#",
	"50d": "M",
	"50n": "M",
}

func GetMeteoIcon(icon string) string {
	val, found := lookupMap[icon]
	if !found {
		return ")"
	}
	return val
}
