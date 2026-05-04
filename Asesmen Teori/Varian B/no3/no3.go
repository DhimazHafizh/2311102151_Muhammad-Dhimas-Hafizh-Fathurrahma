package main

import "fmt"

// function sesuai perintah soal
func tampilKelipatan(n int) {
	if n%2 == 0 {
		// genap → horizontal
		for i := 1; i <= n; i++ {
			fmt.Printf("%d ", i*5)
		}
		fmt.Println()
	} else {
		// ganjil → vertikal
		for i := 1; i <= n; i++ {
			fmt.Println(i * 5)
		}
	}
}

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	// panggil function
	tampilKelipatan(n)
}

/*
misal, input 10 (genap):
5 10 15 20 25 30 35 40 45 50

misal, input 5 (ganjil):
5
10
15
20
25
*/
