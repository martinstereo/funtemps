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
// FAIL returner 0
func FahrenheitToCelsius(fahr float64) float64 {
	celsius := (fahr - 32) * 5 / 9
	return celsius
}

// Kelvin = (Farhenheit - 32) * (5/9) + 273.15
func FahrenheitToKelvin(fahr float64) float64 {
	kelvin := (fahr-32)*5/9 + 273.15
	return kelvin
}

// Fahrenheit = Celsius*(9/5) + 32
func CelsiusToFahrenheit(celsius float64) float64 {
	fahr := celsius*9/5 + 32
	return fahr
}

// Kelvin = Celsius + 273.15
// PASS
func CelsiusToKelvin(celsius float64) float64 {
	kelvin := celsius + 273.15
	return kelvin
}

// Farhrenheit = (Kelvin - 273.15)*(9/5) + 32
func KelvinToFahrenheit(kelvin float64) float64 {
	fahr := (kelvin-273.15)*9/5 + 32
	return fahr
}

// Celsius = Kelvin - 273.15
// PASS
func KelvinToCelsius(kelvin float64) float64 {
	celsius := kelvin - 273.15
	return celsius
}
