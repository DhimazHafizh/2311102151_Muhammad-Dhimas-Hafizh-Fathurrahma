package main

import "fmt"

// function (modularity)
func deretGanjil(n int) {
	if n%2 == 0 {
		// horizontal
		for i := 1; i <= n; i += 2 {
			fmt.Print(i, " ")
		}
	} else {
		// vertikal
		for i := 1; i <= n; i += 2 {
			fmt.Println(i)
		}
	}
}

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	// panggil function
	deretGanjil(n)
}

/*
misal, input 10 (genap):
1 3 5 7 9

misal, input 5 (ganjil):
1
3
5
*/
