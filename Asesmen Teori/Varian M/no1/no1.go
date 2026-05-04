package main

import "fmt"

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	for i := n; i >= 1; i-- {
		fmt.Println(i, "x", i-1, "=", i*(i-1))
	}
}

/*
misal, jika input nya 4,outputnya:
4x3=12
3x2=6
2x1=2
1x0=0

misal, jika input nya 5,outputnya:
5x4=20
4x3=12
3x2=6
2x1=2
1x0=0

*/
