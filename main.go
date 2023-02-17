package main

import (
	"flag"
	"fmt"

	"github.com/naausicaa/funtemps/conv" // importerer konverterings-funksjonene fra conv-mappen
	"github.com/naausicaa/funtemps/format"
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

	flag.Parse() // Parse has to be run after decleration of flags and before their access to the program

	if out == "C" && isFlagPassed("F") && !isFlagPassed("t") { //Fahrenheit til Celsius
		celsius = conv.FahrenheitToCelsius(fahr)
		celsiusFormatted := (format.FormatOutput(celsius)) // Formatted float with no trailing zeros and two decimal places
		fahrFormatted := format.FormatInput(fahr)
		fmt.Printf("%v°F er %v°C\n", fahrFormatted, celsiusFormatted)
	}
	if out == "K" && isFlagPassed("F") && !isFlagPassed("t") { //Fahrenheit til Kelvin
		kelvin = conv.FahrenheitToKelvin(fahr)
		kelvinFormatted := format.FormatOutput(kelvin)
		fahrFormatted := format.FormatInput(fahr)
		fmt.Printf("%v°F er %v°K\n", fahrFormatted, kelvinFormatted)
	}
	if out == "F" && isFlagPassed("C") && !isFlagPassed("t") { //Celsius til Fahrenheit
		fahr = conv.CelsiusToFahrenheit(celsius)
		fahrFormatted := format.FormatOutput(fahr)
		celsiusFormatted := format.FormatInput(celsius)
		fmt.Printf("%v°C er %v°F\n", celsiusFormatted, fahrFormatted)
	}
	if out == "K" && isFlagPassed("C") && !isFlagPassed("t") { //Celsius til Kelvin
		kelvin = conv.CelsiusToKelvin(celsius)
		kelvinFormatted := format.FormatOutput(kelvin)
		celsiusFormatted := format.FormatInput(celsius)
		fmt.Printf("%v°C er %v°K\n", celsiusFormatted, kelvinFormatted)
	}
	if out == "F" && isFlagPassed("K") && !isFlagPassed("t") { //Kelvin til Fahrenheit
		fahr = conv.KelvinToFahrenheit(kelvin)
		fahrFormatted := format.FormatOutput(fahr)
		kelvinFormatted := format.FormatInput(kelvin)
		fmt.Printf("%v°K er %v°F\n", kelvinFormatted, fahrFormatted)
	}
	if out == "C" && isFlagPassed("K") && !isFlagPassed("t") { //Kelvin til Celsius
		celsius = conv.KelvinToCelsius(kelvin)
		celsiusFormatted := format.FormatOutput(celsius)
		kelvinFormatted := format.FormatInput(kelvin)
		fmt.Printf("%v°K er %v°C\n", kelvinFormatted, celsiusFormatted)
	}

	// Funfacts - Sun
	if funf == "sun" && isFlagPassed("funfacts") && isFlagPassed("t") && !isFlagPassed("out") {
		sunFact := funfacts.GetFunFacts(funf)
		if t == "C" { // if user typed Celsius
			fmt.Printf("%v %v°C.\n", sunFact[0], format.FormatOutput(15000000))
			fmt.Printf("%v %v°C.\n", sunFact[1], format.FormatOutput(conv.KelvinToCelsius(5778.00)))
		} else if t == "K" { // if user typed Kelvin
			fmt.Printf("%v %v°K.\n", sunFact[0], format.FormatOutput(conv.CelsiusToKelvin(15000000)))
			fmt.Printf("%v %v°K.\n", sunFact[1], format.FormatOutput(5778))
		} else if t == "F" { // if user typed Fahrenheit
			fmt.Printf("%v %v°F.\n", sunFact[0], format.FormatOutput(conv.CelsiusToFahrenheit(15000000)))
			fmt.Printf("%v %v°F.\n", sunFact[1], format.FormatOutput(conv.KelvinToFahrenheit(5778)))
		}
	}

	// Luna facts
	if funf == "luna" && isFlagPassed("funfacts") && isFlagPassed("t") && !isFlagPassed("out") {
		lunafact := funfacts.GetFunFacts(funf)
		if t == "C" {
			fmt.Printf("%v %v°C.\n", lunafact[0], -183)
			fmt.Printf("%v %v°C.\n", lunafact[1], 106)
		} else if t == "K" {
			fmt.Printf("%v %v°K.\n", lunafact[0], format.FormatOutput(conv.CelsiusToKelvin(-183)))
			fmt.Printf("%v %v°K.\n", lunafact[1], format.FormatOutput(conv.CelsiusToKelvin(106)))
		} else if t == "F" {
			fmt.Printf("%v %v°F.\n", lunafact[0], format.FormatOutput(conv.CelsiusToFahrenheit(-183)))
			fmt.Printf("%v %v°F.\n", lunafact[1], format.FormatOutput(conv.CelsiusToFahrenheit(106)))
		}
	}

	//Terra facts
	if funf == "terra" && isFlagPassed("funfacts") && isFlagPassed("t") && !isFlagPassed("out") {
		terrafact := funfacts.GetFunFacts(funf)
		if t == "C" {
			fmt.Printf("%v %v°C.\n", terrafact[0], 56.7)
			fmt.Printf("%v %v°C.\n", terrafact[1], -89.4)
		} else if t == "K" {
			fmt.Printf("%v %v°K.\n", terrafact[0], 329.82)
			fmt.Printf("%v %v°K.\n", terrafact[1], format.FormatOutput(conv.CelsiusToKelvin(-89.4)))
		} else if t == "F" {
			fmt.Printf("%v %v°F.\n", terrafact[0], 134)
			fmt.Printf("%v %v°F.\n", terrafact[1], format.FormatOutput(conv.CelsiusToFahrenheit(-89.4)))
		}
	}

	// Error message if less than 2 flags were provided
	if flag.NFlag() == 0 { // error if no flags were provided
		fmt.Println("No flags were provided... Try -help ...", flag.NFlag())
	} else if flag.NFlag() == 1 {
		fmt.Println("Only one flag was provided... Program ran with default values ...", flag.NFlag())
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
