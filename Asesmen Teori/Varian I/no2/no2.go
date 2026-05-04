package main

import "fmt"

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	if n%2 == 1 {
		fmt.Println("Inputan adalah ganjil")
	}

	for i := 1; i <= n; i++ {
		fmt.Println(i, "x", i, "=", i*i)
	}

	if n%2 == 0 {
		fmt.Println("Inputan adalah genap")
	}
}

/*
Misal, Jika input ganjil (3):
“Inputan adalah ganjil”
1x1=1
2x2=4
3x3=9

Misal, Jika input genap (4):
1x1=1
2x2=4
3x3=9
4x4=16
“Inputan adalah genap”
*/
