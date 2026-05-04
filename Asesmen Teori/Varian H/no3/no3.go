package main

import "fmt"

// function
func tampil(n int) {
	if n%2 == 1 {
		// ganjil
		for i := 1; i <= n; i++ {
			if i%2 == 1 {
				fmt.Println("ganjil")
			} else {
				fmt.Println(i, ". #")
			}
		}
	} else {
		// genap
		for i := 1; i <= n; i++ {
			if i%2 == 1 {
				fmt.Println(i, ". #")
			} else {
				fmt.Println("genap")
			}
		}
	}
}

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	tampil(n)
}

/*
misal, jika input nya 5 atau angka ganjil,outputnya:
ganjil
2. #
ganjil
4. #
ganjil

misal, jika input nya 10 atau angka genap,outputnya:
1. #
genap
3. #
genap
5. #
genap
7. #
genap
9. #
Genap
*/
