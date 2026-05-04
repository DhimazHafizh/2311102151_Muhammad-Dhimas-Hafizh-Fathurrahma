package main

import "fmt"

// function
func faktorial(n int) {
	fmt.Print("inputan adalah ", n, ", maka faktorialnya ")

	for i := n; i >= 1; i-- {
		fmt.Print(i)
		if i > 1 {
			fmt.Print("x")
		}
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	faktorial(n)
}

/*
Misal, jika Inputan 5, maka outputnya:
"inputan adalah 5, maka faktorialnya 5x4x3x2x1”

Misal, jika Inputan 8, maka outputnya:
"inputan adalah 8, maka faktorialnya 8x7x6x5x4x3x2x1”
*/
