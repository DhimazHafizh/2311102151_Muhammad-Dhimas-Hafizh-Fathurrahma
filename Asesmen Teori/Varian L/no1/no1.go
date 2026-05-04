package main

import "fmt"

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	for i := 10; i > 10-n; i-- {
		fmt.Print(i, " ")
	}
}

/*
misal, jika input nya 10,outputnya:
10 9 8 7 6 5 4 3 2 1

misal, jika input nya 5,outputnya:
10 9 8 7 6

misal, jika input nya 3,outputnya:
10 9 8

*/
