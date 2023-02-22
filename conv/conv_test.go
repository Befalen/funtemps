package conv

import (
	"reflect"
	"testing"
)

/*
*
Mal for testfunksjoner
Du skal skrive alle funksjonene basert på denne malen
For alle konverteringsfunksjonene (tilsammen 6)
kan du bruke malen som den er (du må selvsagt endre
funksjonsnavn og testverdier
*/

func TestFahrenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67},
	}

	for _, tc := range tests {
		got := FahrenheitToCelsius(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

// De andre testfunksjonene implementeres her
// ...

// Konverterer fra Fahrenheit til Celsius
func FahrenheitToCelsius(value float64) float64 {
	// Formellen
	return (value - 32) * (5.0 / 9.0)
}
