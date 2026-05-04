package main

import "fmt"

type bilangan int
type pecahan float64

func main() {
	var a, b bilangan
	var hasil pecahan

	fmt.Print("masukkan bilangan 1 : ")
	fmt.Scan(&a)
	fmt.Print("masukkan bilangan 2 : ")
	fmt.Scan(&b)

	hasil = pecahan(a) / pecahan(b)
	fmt.Println(hasil)
}
