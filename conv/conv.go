package conv

func CelsiusToFahrenheit(value float64) float64 {
	return (value * 9.0 / 5.0) + 32.0
}

func CelsiusToKelvin(value float64) float64 {
	return value + 273.15
}

func FahrenheitToCelsius(value float64) float64 {
	return (value - 32.0) * 5.0 / 9.0
}

func FahrenheitToKelvin(value float64) float64 {
	return (value-32.0)*5.0/9.0 + 273.15
}

func KelvinToCelsius(value float64) float64 {
	return value - 273.15
}

func KelvinToFahrenheit(value float64) float64 {
	return (value-273.15)*9.0/5.0 + 32.0
}
