package funfacts

import (
	"reflect"
	"testing"
)

/*
*

	Mal for TestGetFunFacts funksjonen.
	Definer korrekte typer for input og want,
	og sette korrekte testverdier i slice tests.
*/
func TestGetFunFacts(t *testing.T) {
	type test struct {
		input string
		want  []string
	}

	// Her må du legge inn korrekte testverdier
	tests := []test{
		{
			input: "sun",
			want: []string{
				"Temperaturen i Solens kjerne er",
				"Temperatur på ytre lag av Solen er",
			},
		},

		{
			input: "luna",
			want: []string{
				"Temperatur på Månens overflate om natten er",
				"Temperatur på Månens overflate om dagen er",
			},
		},

		{
			input: "terra",
			want: []string{
				"Høyeste temperatur målt på Jordens overflate er",
				"Laveste temperatur målt på Jordens overflate er",
			},
		},
	}

	for _, tc := range tests {
		got := GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
