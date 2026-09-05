// Package weather provides information about the current
// weather conditions for a given location.
package weather

var (
	// CurrentCondition stores a string describing the current weather condition.
	CurrentCondition string
	// CurrentLocation stores the name of a city for the associated weather condition.
	CurrentLocation string
)

// Forecast returns the current weather conditions for a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
