package conv

func CelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9.0 / 5.0) + 32.0
}

func CelsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

func FahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32.0) * 5.0 / 9.0
}

func FahrenheitToKelvin(fahrenheit float64) float64 {
	return (fahrenheit-32.0)*5.0/9.0 + 273.15
}

func KelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

func KelvinToFahrenheit(kelvin float64) float64 {
	return (kelvin-273.15)*9.0/5.0 + 32.0
}
