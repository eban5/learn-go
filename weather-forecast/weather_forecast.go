// Package weather provides the forecast for a given location and current conditions.
package weather

// CurrentCondition holds a string with the current conditions.
var CurrentCondition string

// CurrentLocation holds a string with the current location.
var CurrentLocation string

// Forecast returns the conditions for the current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
