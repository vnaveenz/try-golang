package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type

func (tu TemperatureUnit) String() string {
	switch tu {
	case Celsius:
		return "°C"
	case Fahrenheit:
		return "°F"
	default:
		return "Invalid Temperature Unit"
	}
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type

func (temp Temperature) String() string {
	return fmt.Sprintf("%v %s", temp.degree, temp.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit

func(su SpeedUnit) String() string {
	switch su {
	case KmPerHour:
		return "km/h"
	case MilesPerHour:
		return "mph"
	default:
		return "Invalid speed unit"
	}
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed

func (sp Speed) String() string {
	return fmt.Sprintf("%v %s", sp.magnitude, sp.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData

func(md MeteorologyData) String() string {
	return fmt.Sprintf("%s: %v %s, Wind %s at %v %s, %v%% Humidity", md.location, md.temperature.degree,md.temperature.unit, md.windDirection , md.windSpeed.magnitude,md.windSpeed.unit,md.humidity)
}
