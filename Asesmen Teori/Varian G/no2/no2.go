package main

import "fmt"

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	// deret 0+1+2+...
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print("+")
		}
		fmt.Print(i)
	}
	fmt.Println()

	// keterangan
	if n%2 == 0 {
		fmt.Println("inputan adalah angka genap")
	} else {
		fmt.Println("inputan adalah angka ganjil")
	}
}

/*
misal, jika input nya 5 atau angka ganjil,outputnya:
0+1+2+3+4
inputan adalah angka ganjil

misal, jika input nya 10 atau angka genap,outputnya:
0+1+2+3+4+5+6+7+8+9
inputan adalah angka genap

*/
