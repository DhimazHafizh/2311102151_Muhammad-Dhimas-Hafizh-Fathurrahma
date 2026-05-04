package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n%2 == 1 {
		for i := 1; i <= n; i++ {
			fmt.Print("## ")
		}
		fmt.Println("Inputan adalah Angka Ganjil")
	} else {
		for i := 1; i <= n; i++ {
			fmt.Print("#- ")
		}
		fmt.Println("Inputan adalah Angka Genap")
	}
}

/*
misal, input 5 (ganjil):
## ## ## ## ##
Inputan adalah Angka Ganjil

misal, input 4 (genap):
#- #- #- #-
Inputan adalah Angka Genap

*/
