// Package weather reprents this go package.
package weather

// CurrentCondition represents the current Weather condition.
var CurrentCondition string

// CurrentLocation represents the name of current City.
var CurrentLocation string

// Forecast tells us current weather condition of requested City.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
