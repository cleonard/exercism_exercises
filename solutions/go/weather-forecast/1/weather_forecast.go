// Package weather forecasts weather in various cities.
package weather

var (
	// CurrentCondition refers to the current condition.
	CurrentCondition string
	// CurrentLocation refers to the current location.
	CurrentLocation string
)

// Forecast returnsa statement about the current weather conditions in the current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
