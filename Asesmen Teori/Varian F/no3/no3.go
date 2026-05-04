package main

import "fmt"

func tampil(n int) {
	if n%2 == 0 {
		for i := -n; i <= 0; i++ {
			fmt.Print(i, " ")
		}
	} else {
		for i := n; i >= 0; i-- {
			fmt.Println("-", i, i)
		}
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	tampil(n)
}

/*
misal, jika input nya 10 atau angka genap:
-10 -9 -8 -7 -6 -5 -4 -3 -2 -1 -0

misal, jika input nya 5 atau angka ganjil:
-55
-44
-33
-22
-11
-00
*/
