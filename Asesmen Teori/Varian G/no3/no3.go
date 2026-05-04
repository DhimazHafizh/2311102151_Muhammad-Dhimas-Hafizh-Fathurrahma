package main

import "fmt"

// function
func kuadrat(n int) {
	fmt.Println("inputan adalah", n, ", maka", n, "*", n, "=", n*n)
}

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	kuadrat(n)
}

/*
misal, jika input nya 10,outputnya:
inputan adalah 10,  maka 10*10=100

misal, jika input nya 7,outputnya:
inputan adalah 7,  maka 7*7=49

*/
