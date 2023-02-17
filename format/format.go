package format

import (
	"fmt"
	"strings"
)

// Function that format numbers to no trailing zeros and two decimals.
// Also adding spaces between big numbers, works for integers as well.
// From Chat GPT

func FormatOutput(num float64) string {
	str := fmt.Sprintf("%.2f", num)
	str = strings.TrimRight(str, "0")
	parts := strings.Split(str, ".")
	integerPart := parts[0]
	decimalPart := parts[1]

	var formattedIntegerPart string
	n := len(integerPart)
	for i, v := range integerPart {
		formattedIntegerPart += string(v)
		if (n-i-1)%3 == 0 && i != n-1 {
			formattedIntegerPart += " "
		}
	}
	if decimalPart != "" {
		return formattedIntegerPart + "." + decimalPart
	}
	return formattedIntegerPart
}

// Veldig Lik funksjon men forskjellen er at for output skal den ikke formattere til 2 desimaler.
// men heller representere tallet brukeren brukeren skrev inn på den mest treffsikre måten
func FormatInput(num float64) string {
	str := fmt.Sprintf("%f", num)
	str = strings.TrimRight(str, "0")
	parts := strings.Split(str, ".")
	integerPart := parts[0]
	decimalPart := parts[1]

	var formattedIntegerPart string
	n := len(integerPart)
	for i, v := range integerPart {
		formattedIntegerPart += string(v)
		if (n-i-1)%3 == 0 && i != n-1 {
			formattedIntegerPart += " "
		}
	}
	if decimalPart != "" {
		return formattedIntegerPart + "." + decimalPart
	}
	return formattedIntegerPart
}
