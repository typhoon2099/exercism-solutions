package meteorology

import "fmt"

type TemperatureUnit int

func (t *TemperatureUnit) String() string {
    if *t == Celsius {
        return "°C"
    }

    return "°F"
}

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (t *Temperature) String() string {
    return fmt.Sprintf("%d %s", t.degree, t.unit.String())
}

type SpeedUnit int

func (t *SpeedUnit) String() string {
    if *t == KmPerHour {
        return "km/h"
    }

    return "mph"
}

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (t *Speed) String() string {
    return fmt.Sprintf("%d %s", t.magnitude, t.unit.String())
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (t *MeteorologyData) String() string {
    return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity", t.location, t.temperature.String(), t.windDirection, t.windSpeed.String(), t.humidity)
}
