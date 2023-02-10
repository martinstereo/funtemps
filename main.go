package main

import (
	"flag"
	"fmt"

	"github.com/naausicaa/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var (
	fahr     float64
	kelvin   float64
	celsius  float64
	out      string
	funfacts string
)

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.

func init() {

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader Fahrenheit")
	flag.Float64Var(&kelvin, "K", 0.0, "temperatur i grader Kelvin")
	flag.Float64Var(&celsius, "C", 0.0, "temperatur i grader Celsius")

	flag.StringVar(&out, "out", "C", "beregne temperatur i F - farhenheit")

	// flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")

	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises
}

func main() {
	fmt.Println("Starting the application...")
	flag.Parse()

	/**
	    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
	    pakkene implementere
	    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
	    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
	    hvor mange flagg og argumenter er spesifisert på kommandolinje.

	        fmt.Println("len(flag.Args())", len(flag.Args()))
			    fmt.Println("flag.NFlag()", flag.NFlag())

	    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
	    brukes for å utelukke ugyldige kombinasjoner:
	    -F, -C, -K kan ikke brukes samtidig
	    disse tre kan brukes med -out, men ikke med -funfacts
	    -funfacts kan brukes kun med -t
	    ...
	*/

	//Fahrenheit til Celsius
	if out == "C" && isFlagPassed("F") {
		celsius = conv.FahrenheitToCelsius(fahr)
		fmt.Println(fahr, "°F er ", celsius, "°C")
	}

	//Fahrenheit til Kelvin
	if out == "K" && isFlagPassed("F") {
		kelvin = conv.FahrenheitToKelvin(fahr)
		fmt.Println(fahr, "°F er ", kelvin, "°K")
	}

	//Celsius til Fahrenheit
	if out == "F" && isFlagPassed("C") {
		fahr = conv.CelsiusToFahrenheit(celsius)
		fmt.Println(celsius, "°C er ", fahr, "°F")
	}

	//Celsius til Kelvin
	if out == "K" && isFlagPassed("C") {
		kelvin = conv.CelsiusToKelvin(celsius)
		fmt.Println(celsius, "°C er ", kelvin, "°K")
	}

	//Kelvin til Fahrenheit
	if out == "F" && isFlagPassed("K") {
		fahr = conv.KelvinToFahrenheit(kelvin)
		fmt.Println(kelvin, "°K er ", fahr, "°F")
	}

	//Kelvin til Celsius
	if out == "C" && isFlagPassed("K") {
		celsius = conv.KelvinToCelsius(kelvin)
		fmt.Println(kelvin, "°K er ", celsius, "°C")
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
