package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

/*
// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(value float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
	return 56.67
	}
*/

// fahr til cel
func FahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

// cel til fahr
func CelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

// kel til cel
func KelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

// cel til kel
func CelsiusToKelvin(value float64) float64 {
	return (value - 273.15)
}

// farh til kel
func FarhenheitToKelvin(value float64) float64 {
	return (value - 32*(5/9) + 273.15)
}

// kel til fahr
func KelvinToFarhenheit(value float64) float64 {
	return (value - 273.15*(9/5) + 32)
}
