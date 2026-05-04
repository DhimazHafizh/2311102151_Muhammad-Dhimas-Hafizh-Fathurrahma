package main

import "fmt"

// function (modular)
func polaSimbol(n int) {
	if n%2 == 1 {
		// ganjil → pakai label
		for i := 1; i <= n; i++ {
			if i%2 == 1 {
				fmt.Println(i, ". * (Ganjil)")
			} else {
				fmt.Println(i, ". * (Genap)")
			}
		}
	} else {
		// genap → tanpa label
		for i := 1; i <= n; i++ {
			fmt.Println(i, ". *")
		}
	}
}

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	// panggil function
	polaSimbol(n)
}

/*
misal, inputan 5 (ganjil):
1. * (Ganjil)
2. * (Genap)
3. * (Ganjil)
4. * (Genap)
5. * (Ganjil)

misal, inputan 4 (genap):
1. *
2. *
3. *
4. *

*/
