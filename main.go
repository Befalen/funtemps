package main

import (
	"flag"
	"fmt"

	"github.com/befalen/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var kel float64
var cel float64
var out string
var funfacts string

func init() {

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i fahrenheit")
	flag.Float64Var(&kel, "K", 0.0, "temperatur i kelvin")
	flag.Float64Var(&cel, "C", 0.0, "temperatur i celsius")

	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celcius, F - fahrenheit, K - Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
}

func main() {
	flag.Parse()

	// kel til cel
	if isFlagPassed("K") && out == "C" {
		output := conv.KelvinToCelsius(kel)
		fmt.Printf("%v K er %v C", kel, output)
	}
	// cel til kel
	if isFlagPassed("C") && out == "K" {
		output := conv.CelsiusToKelvin(cel)
		fmt.Printf("%v C er %v K", kel, output)
	}

	// fahr til cel
	if isFlagPassed("F") && out == "C" {
		output := conv.FahrenheitToCelsius(fahr)
		fmt.Printf("%v F er %v C", kel, output)
	}
	// cel til fahr
	if isFlagPassed("C") && out == "F" {
		output := conv.CelsiusToFahrenheit(cel)
		fmt.Printf("%v C er %v F", kel, output)
	}

	// fahr til kel
	if isFlagPassed("F") && out == "K" {
		output := conv.FahrenheitToKelvin(fahr)
		fmt.Printf("%v F er %v K", kel, output)
	}
	// kel til fahr
	if isFlagPassed("K") && out == "F" {
		output := conv.KelvinToFahrenheit(kel)
		fmt.Printf("%v K er %v F", kel, output)
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
