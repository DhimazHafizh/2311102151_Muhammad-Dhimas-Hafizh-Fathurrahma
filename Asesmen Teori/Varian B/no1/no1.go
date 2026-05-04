package main

import "fmt"

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Println(i * 5)
	}
}

/* 
misal, jika input nya 10:
5
10
15
20
25
30
35
40
45
50

*/