package main

import "fmt"

func cetakDeret(n int) {
	fmt.Print(n, " ")
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
			fmt.Print(n, " ")
		} else if n%2 != 0 {
			n = 3*n + 1
			fmt.Print(n, " ")
		} else if n == 1 {
			fmt.Print(n)
		}
	}
}

func main() {
	var bilangan int

	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&bilangan)

	cetakDeret(bilangan)
}
