package main

import (
	"flag"
	"fmt"

	"github.com/Befalen/funtemps/conv"
	"github.com/Befalen/funtemps/funfacts"
)

// Definerer flag-variablene i hoved-"scope"
var value float64
var unit string
var out string
var funfactsFlag string
var temperatureScale string

func init() {

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&value, "v", 0.0, "temperature value")
	flag.StringVar(&unit, "u", "C", "temperature unit, e.g. C, F, K")

	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfactsFlag, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	flag.StringVar(&temperatureScale, "t", "C", "temperaturskala som skal brukes når fun-facts vises")

}

func main() {

	flag.Parse()

	// Check for conflicting flags
	if isFlagPassed("v") && isFlagPassed("funfacts") || isFlagPassed("t") {
		fmt.Println("Error: -v kan ikke bli brukt med -funfacts eller -t flags")
		return
	}

	if isFlagPassed("out") {
		switch out {
		case "C":
			result := conv.Convert(value, unit, "C")
			fmt.Printf("%0.2f%s is %0.2f\n", value, unit, result)
		case "F":

			result := conv.Convert(value, unit, "F")
			fmt.Printf("%0.2f%s is %0.2f\n", value, unit, result)
		case "K":

			result := conv.Convert(value, unit, "K")
			fmt.Printf("%0.2f%s is %0.2f\n", value, unit, result)
		default:
			fmt.Println("Ugyldig utdataflagg. Mulige verdier er C, F, K")
		}
	}

	if isFlagPassed("funfacts") {
		funFacts := funfacts.GetFunFacts(funfactsFlag)
		for _, fact := range funFacts {
			fmt.Println(fact)
		}
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
