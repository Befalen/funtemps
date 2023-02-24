package main

import (
	"github.com/Befalen/funtemps/conv"
)

func temperatureConv(input float64, inputUnit string, outputUnit string) float64 {
	switch inputUnit {
	case "C":
		switch outputUnit {
		case "C":
			return input
		case "F":
			return conv.CelsiusToFahrenheit(input)
		case "K":
			return conv.CelsiusToKelvin(input)
		}
	case "F":
		switch outputUnit {
		case "C":
			return conv.FahrenheitToCelsius(input)
		case "F":
			return input
		case "K":
			return conv.FahrenheitToKelvin(input)
		}
	case "K":
		switch outputUnit {
		case "C":
			return conv.KelvinToCelsius(input)
		case "F":
			return conv.KelvinToFahrenheit(input)
		case "K":
			return input
		}
	}
	return input
}
