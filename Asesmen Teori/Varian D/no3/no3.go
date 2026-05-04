package main

import "fmt"

// function
func simbol(n int) {
	if n%2 == 1 {
		for i := 1; i <= n; i++ {
			fmt.Print("## ")
		}
		fmt.Println("\nInputan adalah Angka Ganjil")
	} else {
		for i := 1; i <= n; i++ {
			fmt.Print("#- ")
		}
		fmt.Println("\nInputan adalah Angka Genap")
	}
}

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	// panggil function
	simbol(n)
}

/*
misal, input 5 (ganjil):
## ## ## ## ##
Inputan adalah Angka Ganjil

misal, input 4 (genap):
#- #- #- #-
Inputan adalah Angka Genap
*/
