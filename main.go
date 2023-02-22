package main

import (
	"flag"
	"fmt"
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
	if isFlagPassed("v") && (isFlagPassed("funfacts") || isFlagPassed("t")) {
		fmt.Println("Error: -v kan ikke bli brukt med -funfacts eller -t flags")
		return
	}

	if isFlagPassed("out") {
		var result float64
		switch out {
		case "C":
			result = convert(value, unit, "C")
		case "F":
			result = convert(value, unit, "F")
		case "K":
			result = convert(value, unit, "K")
		default:
			fmt.Println("Ugyldig utdataflagg. Mulige verdier er C, F, K")
			return
		}
		fmt.Printf("%0.2f%s is %0.2f%s\n", value, unit, result, out)
	}

	if isFlagPassed("funfacts") {
		funFacts := getFunFacts(funfactsFlag)
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

func getFunFacts(subject string) []string {
	var facts []string
	switch subject {
	case "sun":
		facts = []string{
			"The sun is the closest star to Earth.",
		}
	case "moon":
		facts = []string{
			"The moon is about 4.5 billion years old.",
		}
	case "earth":
		facts = []string{
			"Earth is the third planet from the sun.",
		}
	default:
		facts = []string{"Sorry, I don't have any fun facts for you."}
	}
	return facts
}
