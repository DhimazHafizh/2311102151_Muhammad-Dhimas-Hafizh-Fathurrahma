package main

import "fmt"

func penjumlahan(n int) int {
	if n < 0 {
		return 0
	} else {
		return n + penjumlahan(n-1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan input : ")
	fmt.Scan(&n)
	fmt.Println(penjumlahan(n))
}

// penjumlahan(5)
// = 5 + penjumlahan(4)
// = 5 + (4 + penjumlahan(3))
// = 5 + 4 + (3 + penjumlahan(2))
// = 5 + 4 + 3 + (2 + penjumlahan(1))
// = 5 + 4 + 3 + 2 + 1
// = 15
