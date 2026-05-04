package main

import "fmt"

// function
func perkalian(n int) {
	if n%2 == 0 {
		for i := n; i >= 1; i-- {
			if i%2 == 0 {
				fmt.Println(0)
			} else {
				fmt.Println(i, "x", i-1, "=", i*(i-1))
			}
		}
	} else {
		for i := n; i >= 1; i-- {
			if i%2 == 1 {
				fmt.Println(0)
			} else {
				fmt.Println(i, "x", i-1, "=", i*(i-1))
			}
		}
	}
}

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	// panggil function
	perkalian(n)
}

/*
misal, jika input nya 4 (genap),outputnya:
0
3x2=6
0
1x0=0

misal, jika input nya 5 (ganjil),outputnya:
0
4x3=12
0
2x1=2
0

*/
