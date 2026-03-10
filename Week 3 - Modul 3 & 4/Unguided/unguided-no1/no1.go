package main

import "fmt"

func main() {
	var tahun int
	var Kabisat bool = false

	fmt.Print("Masukkan tahun : ")
	fmt.Scan(&tahun)

	if tahun % 4 == 0 || tahun % 400 == 0 {
		Kabisat = true
	} else {
		Kabisat = false
	}

	fmt.Println("Kabisat :", Kabisat)
}
