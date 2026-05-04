package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n%2 == 0 {
		for i := 1; i <= n; i += 2 {
			fmt.Print(i, " ")
		}
	} else {
		for i := 1; i <= n; i += 2 {
			fmt.Println(i)
		}
	}
}

/*
misal, input 10 (genap):
1 3 5 7 9

misal, input 5 (ganjil):
1
3
5

*/
