package conv

/*
  Konverteringsformlene for Celcius, Kelvin og Farhenheit:

  * Celsius = Kelvin - 273.15
  * Kelvin = Celsius + 273.15
  * Fahrenheit = Celsius*(9/5) + 32
  * Celsius = (Farhrenheit - 32)*(5/9)
  * Farhrenheit = (Kelvin - 273.15)*(9/5) + 32
  * Kelvin = (Farhenheit - 32) * (5/9) + 273.15

*/

// Celsius = (Farhrenheit - 32)*5/9
func FahrenheitToCelsius(fahr float64) float64 {
	return (fahr - 32) * 5 / 9
}

// Kelvin = (Farhenheit - 32) * (5/9) + 273.15
func FahrenheitToKelvin(fahr float64) float64 {
	return (fahr-32)*5/9 + 273.15
}

// Fahrenheit = Celsius*(9/5) + 32
func CelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

// Kelvin = Celsius + 273.15
func CelsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

// Farhrenheit = (Kelvin - 273.15)*(9/5) + 32
func KelvinToFahrenheit(kelvin float64) float64 {
	return (kelvin-273.15)*9/5 + 32
}

// Celsius = Kelvin - 273.15
func KelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}
