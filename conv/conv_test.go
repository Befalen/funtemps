package conv

import (
	"reflect"
	"testing"
)

// fahr til cel
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

// cel til fahr
func TestCelsiusToDFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}
	tests := []test{
		{input: 56.67, want: 134},
	}
	for _, tc := range tests {
		got := FahrenheitToCelsius(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

// kel til cel
func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}
	tests := []test{
		{input: 329.82, want: 56.67},
	}

	for _, tc := range tests {
		got := FahrenheitToCelsius(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

// cel til kel
func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}
	tests := []test{
		{input: 56.67, want: 329.82},
	}

	for _, tc := range tests {
		got := FahrenheitToCelsius(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

// fahr til kel
func TestFahrenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}
	tests := []test{
		{input: 134, want: 329.82},
	}

	for _, tc := range tests {
		got := FahrenheitToCelsius(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

// kel til fahr
func TestKelvinToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}
	tests := []test{
		{input: 329.82, want: 134},
	}

	for _, tc := range tests {
		got := FahrenheitToCelsius(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
