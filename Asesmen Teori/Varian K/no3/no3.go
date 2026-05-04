package main

import "fmt"

// function
func simbol(n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			fmt.Print("* ")
		} else {
			fmt.Print("@ ")
		}
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
misal, jika input nya 10,outputnya:
@ @ *  @ @ *  @ @ *  @

misal, jika input nya 6, outputnya:
@ @ *  @ @ *
*/
