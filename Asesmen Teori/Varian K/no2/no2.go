package main

import "fmt"

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			fmt.Print("* ")
		} else {
			fmt.Print("@ ")
		}
	}
}

/*
misal, jika input nya 10,outputnya:
@ @ *  @ @ *  @ @ *  @

misal, jika input nya 6, outputnya:
@ @ *  @ @ *
*/
