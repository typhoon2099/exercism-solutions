// Package weather provides weather forecasts for a given city and condition.
package weather

// CurrentCondition is the current weather condition.
var CurrentCondition string
// CurrentLocation is the location you want weather information for.
var CurrentLocation string

// Forecast returns the weather forecast as a String using the current location and current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
