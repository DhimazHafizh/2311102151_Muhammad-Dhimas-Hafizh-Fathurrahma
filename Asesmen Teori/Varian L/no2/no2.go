package main

import "fmt"

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	for i := 10; i > 10-n; i-- {
		if i%2 == 0 {
			fmt.Print("() ")
		} else {
			fmt.Print(i, " ")
		}
	}
}

/*
misal, jika input nya 10,outputnya:
() 9 () 7 () 5 () 3 () 1

misal, jika input nya 6, outputnya:
() 9 () 7 () 5

*/
