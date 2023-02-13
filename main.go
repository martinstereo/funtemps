package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/naausicaa/funtemps/conv"     // importerer konverterings-funksjonene fra conv-mappen
	"github.com/naausicaa/funtemps/funfacts" // importerer funfacts-funksjonen fra funfacts-mappen
)

// Definerer flag-variablene i hoved-"scope"
var (
	fahr    float64
	kelvin  float64
	celsius float64
	out     string
	funf    string // funfacts
	t       string
)

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader Fahrenheit")
	flag.Float64Var(&kelvin, "K", 0.0, "temperatur i grader Kelvin")
	flag.Float64Var(&celsius, "C", 0.0, "temperatur i grader Celsius")
	flag.StringVar(&out, "out", "C", "beregne temperatur i F - farhenheit, C - Celsius eller K - Kelvin")
	flag.StringVar(&funf, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	flag.StringVar(&t, "t", "C", "bestemmer hvilken temperaturskala som skal brukes når funfacts skal vises")
}

/*
Program som kan tar innput i form av (kun en om gangen) temperaturene Celsius, Fahrenheit og Kalvin
fra kommandolinje i form av flagg
F.eks.

	"-F 100 -out C"
	vil returnere 100 grader Fahrenheit til Celsius.

Og returnerer temperaturen spesifisert av input i "-out X"
Kan også returnere funfacts om temperatur på solen, månen og jordkloden ved:

	"-funfacts sun -t C"
	Vil returnere funfacts om sola i celsius, fungerer også for Kelvin og Fahrenheit
*/
func main() {
	fmt.Println("Starting the application...") // Test print

	flag.Parse() // Parse has to be run after decleration of flags and before their access to the program

	if out == "C" && isFlagPassed("F") && !isFlagPassed("t") { //Fahrenheit til Celsius
		celsius = conv.FahrenheitToCelsius(fahr)
		celsiusFormatted := (FormatNumber(celsius)) // Formatted float with no trailing zeros and two decimal places
		fmt.Printf("%f°F er %v°C\n", fahr, celsiusFormatted)
	}
	if out == "K" && isFlagPassed("F") && !isFlagPassed("t") { //Fahrenheit til Kelvin
		kelvin = conv.FahrenheitToKelvin(fahr)
		kelvinFormatted := FormatNumber(kelvin)
		fmt.Printf("%f°F er %v°K\n", fahr, kelvinFormatted)
	}
	if out == "F" && isFlagPassed("C") && !isFlagPassed("t") { //Celsius til Fahrenheit
		fahr = conv.CelsiusToFahrenheit(celsius)
		fahrFormatted := FormatNumber(fahr)
		fmt.Printf("%f°C er %v°F\n", celsius, fahrFormatted)
	}
	if out == "K" && isFlagPassed("C") && !isFlagPassed("t") { //Celsius til Kelvin
		kelvin = conv.CelsiusToKelvin(celsius)
		kelvinFormatted := FormatNumber(kelvin)
		fmt.Printf("%f°C er %v°K\n", celsius, kelvinFormatted)
	}
	if out == "F" && isFlagPassed("K") && !isFlagPassed("t") { //Kelvin til Fahrenheit
		fahr = conv.KelvinToFahrenheit(kelvin)
		fahrFormatted := FormatNumber(fahr)
		fmt.Printf("%f°K er %v°F\n", kelvin, fahrFormatted)
	}
	if out == "C" && isFlagPassed("K") && !isFlagPassed("t") { //Kelvin til Celsius
		celsius = conv.KelvinToCelsius(kelvin)
		celsiusFormatted := FormatNumber(celsius)
		fmt.Printf("%f°K er %v°C\n", kelvin, celsiusFormatted)
	}

	// Funfacts - Sun
	if funf == "sun" && isFlagPassed("funfacts") && isFlagPassed("t") {
		sunFact := funfacts.GetFunFacts(funf)
		if t == "C" { // if user typed Celsius
			fmt.Printf("%v %v°C.\n", sunFact[0], FormatNumber(15000000)) // Skal bytte til printf
			fmt.Printf("%v %v°C.\n", sunFact[1], FormatNumber(conv.KelvinToCelsius(5778.00)))
		} else if t == "K" { // if user typed Kelvin
			fmt.Printf("%v %v°K.\n", sunFact[0], FormatNumber(conv.CelsiusToKelvin(15000000)))
			fmt.Printf("%v %v°K.\n", sunFact[1], FormatNumber(5778))
		} else if t == "F" { // if user typed Fahrenheit
			fmt.Printf("%v %v°F.\n", sunFact[0], FormatNumber(conv.CelsiusToFahrenheit(15000000)))
			fmt.Printf("%v %v°F.\n", sunFact[1], FormatNumber(conv.KelvinToFahrenheit(5778)))
		}
	}

	// Luna facts
	if funf == "luna" && isFlagPassed("funfacts") {
		lunafact := funfacts.GetFunFacts(funf)
		if t == "C" {
			fmt.Printf("%v %v°C.\n", lunafact[0], -183)
			fmt.Printf("%v %v°C.\n", lunafact[1], 106)
		} else if t == "K" {
			fmt.Printf("%v %v°K.\n", lunafact[0], FormatNumber(conv.CelsiusToKelvin(-183)))
			fmt.Printf("%v %v°K.\n", lunafact[1], FormatNumber(conv.CelsiusToKelvin(106)))
		} else if t == "F" {
			fmt.Printf("%v %v°F.\n", lunafact[0], FormatNumber(conv.CelsiusToFahrenheit(-183)))
			fmt.Printf("%v %v°F.\n", lunafact[1], FormatNumber(conv.CelsiusToFahrenheit(106)))
		}
	}

	//Terra facts
	if funf == "terra" && isFlagPassed("funfacts") {
		terrafact := funfacts.GetFunFacts(funf)
		if t == "C" {
			fmt.Printf("%v %v°C.\n", terrafact[0], 56.7)
			fmt.Printf("%v %v°C.\n", terrafact[1], -89.4)
		} else if t == "K" {
			fmt.Printf("%v %v°K.\n", terrafact[0], 329.82)
			fmt.Printf("%v %v°K.\n", terrafact[1], FormatNumber(conv.CelsiusToKelvin(-89.4)))
		} else if t == "F" {
			fmt.Printf("%v %v°F.\n", terrafact[0], 134)
			fmt.Printf("%v %v°F.\n", terrafact[1], FormatNumber(conv.CelsiusToFahrenheit(-89.4)))
		}
	}
	if flag.NFlag() == 0 { // error if no flags were provided
		fmt.Println("No flags were provided... Try -help", flag.NFlag())
	}
}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

// Function that format floats to no trailing zeros and two decimals.
// Also adding spaces between big numbers, works for integers as well.
// From Chat GPT
func FormatNumber(num float64) string {
	str := fmt.Sprintf("%.2f", num)
	str = strings.TrimRight(str, "0")
	str = strings.TrimRight(str, ".")

	parts := strings.Split(str, ".")
	integerPart := parts[0]
	var decimalPart string
	if len(parts) > 1 {
		decimalPart = parts[1]
	}
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
