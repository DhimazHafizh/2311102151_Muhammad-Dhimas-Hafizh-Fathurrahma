package main

import "fmt"

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Println(i, "x", i, "=", i*i)
	}
}

/*
Misal, jika Input 3:
1x1=1
2x2=4
3x3=9

Misal, jika Input 4:
1x1=1
2x2=4
3x3=9
4x4=16

*/
