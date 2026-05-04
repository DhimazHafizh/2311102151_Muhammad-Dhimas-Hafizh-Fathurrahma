package main

import "fmt"

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	if n%2 == 0 {
		// genap
		for i := n; i >= 1; i-- {
			if i%2 == 0 {
				fmt.Println(0)
			} else {
				fmt.Println(i, "x", i-1, "=", i*(i-1))
			}
		}
	} else {
		// ganjil
		for i := n; i >= 1; i-- {
			if i%2 == 1 {
				fmt.Println(0)
			} else {
				fmt.Println(i, "x", i-1, "=", i*(i-1))
			}
		}
	}
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
