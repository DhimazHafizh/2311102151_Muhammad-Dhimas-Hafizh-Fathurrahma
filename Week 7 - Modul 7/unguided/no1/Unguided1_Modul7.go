package main

import "fmt"

type suhu float64

func CelciusToReamur(celcius suhu) suhu {
	return (4.0 / 5.0) * celcius
}

func CelciusToFahrenheit(celcius suhu) suhu {
	return (9.0 / 5.0 * celcius) + 32
}

func CelciusToKelvin(celcius suhu) suhu {
	return celcius + 273.15
}

func main() {
	var suhuInput suhu

	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukkan suhu (celcius) : ")
	fmt.Scan(&suhuInput)

	fmt.Println()
	fmt.Println(suhuInput, "celcius =", CelciusToReamur(suhuInput), "reamur")
	fmt.Println(suhuInput, "celcius =", CelciusToFahrenheit(suhuInput), "fahrenheit")
	fmt.Println(suhuInput, "celcius =", CelciusToKelvin(suhuInput), "kelvin")
}
