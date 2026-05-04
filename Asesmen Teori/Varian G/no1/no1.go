package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print("+")
		}
		fmt.Print(i)
	}
}

/*
misal, jika input nya 10,outputnya:
0+1+2+3+4+5+6+7+8+9

misal, jika input nya 5,outputnya:
0+1+2+3+4

*/
