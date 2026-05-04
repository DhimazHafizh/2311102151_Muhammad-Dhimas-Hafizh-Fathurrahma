package main

import "fmt"

func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * faktorial(n-1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan input : ")
	fmt.Scan(&n)
	fmt.Println(faktorial(n))
}

// faktorial(5)
// = 5 × faktorial(4)
// = 5 × (4 × faktorial(3))
// = 5 × 4 × (3 × faktorial(2))
// = 5 × 4 × 3 × (2 × faktorial(1))
// = 5 × 4 × 3 × 2 × 1
// = 120
