package main

import "fmt"

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	if n%2 == 1 {
		// jika input ganjil
		for i := 1; i <= n; i++ {
			if i%2 == 1 {
				fmt.Println("ganjil")
			} else {
				fmt.Println(i, ". #")
			}
		}
	} else {
		// jika input genap
		for i := 1; i <= n; i++ {
			if i%2 == 1 {
				fmt.Println(i, ". #")
			} else {
				fmt.Println("genap")
			}
		}
	}
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
