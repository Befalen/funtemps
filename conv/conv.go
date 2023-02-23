package conv

// fahr til cel
func FahrenheitToCelsius(fahrenheit float64) float64 {
	return (value - 32) * 5 / 9
}

// cel til fahr
func CelsiusToFahrenheit(celsius float64) float64 {
	return value*9/5 + 32
}

// kel til cel
func KelvinToCelsius(kelvin float64) float64 {
	return value - 273.15
}

// cel til kel
func CelsiusToKelvin(value float64) float64 {
	return value + 273.15
}

// farh til kel
func FahrenheitToKelvin(value float64) float64 {
	return (value - 32*(5/9) + 273.15)
}

// kel til fahr
func KelvinToFahrenheit(value float64) float64 {
	return (value - 273.15*(9/5) + 32)
}
