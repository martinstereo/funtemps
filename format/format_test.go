package format

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
	funksjonsnavn og testverdier)
*/

// HER KOMMER TESTENE

// PASS
func TestFormatInput(t *testing.T) {
	type test struct {
		input float64
		want  string
	}

	tests := []test{
		{input: 5000.00300, want: "5 000.003"},
	}

	for _, tc := range tests {
		got := FormatInput(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

// PASS
func TestFormatOutput(t *testing.T) {
	type test struct {
		input float64
		want  string
	}

	tests := []test{
		{input: 5000.00300, want: "5 000"},
	}

	for _, tc := range tests {
		got := FormatOutput(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
