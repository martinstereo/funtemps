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

/*
Program som kan tar innput i form av (kun en om gangen) temperaturene Celsius, Fahrenheit og Kalvin
fra kommandolinje i form av flagg

F.eks.

"-F 100 -out C"

	vil returnere 100 grader Fahrenheit til Celsius.

Og returnerer temperaturen spesifisert av input i "-out X"
*/

func main() {
	fmt.Println("Starting the application...") // Test print

	flag.Parse() // Parse has to be run after decleration of flags and before their access to the program/main

	/*
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
		fmt.Printf("%.2f°F er %.2f°C", fahr, celsius) // printer ut verdi med 2 desimaltall
	}

	//Fahrenheit til Kelvin
	if out == "K" && isFlagPassed("F") {
		kelvin = conv.FahrenheitToKelvin(fahr)
		fmt.Printf("%.2f °F er %.2f °K", fahr, kelvin)
	}

	//Celsius til Fahrenheit
	if out == "F" && isFlagPassed("C") {
		fahr = conv.CelsiusToFahrenheit(celsius)
		fmt.Printf("%.2f°C er %.2f°F", celsius, fahr)
	}

	//Celsius til Kelvin
	if out == "K" && isFlagPassed("C") {
		kelvin = conv.CelsiusToKelvin(celsius)
		fmt.Printf("%.2f°C er %.2f°K", celsius, kelvin)
	}

	//Kelvin til Fahrenheit
	if out == "F" && isFlagPassed("K") {
		fahr = conv.KelvinToFahrenheit(kelvin)
		fmt.Printf("%.2f°K er %.2f°F", kelvin, fahr)
	}

	//Kelvin til Celsius
	if out == "C" && isFlagPassed("K") {
		celsius = conv.KelvinToCelsius(kelvin)
		fmt.Printf("%.2f°K er %.2f°C", kelvin, celsius)
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
