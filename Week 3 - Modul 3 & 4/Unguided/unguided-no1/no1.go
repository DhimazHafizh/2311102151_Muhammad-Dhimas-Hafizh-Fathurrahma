package main

import "fmt"

func faktorial(n int) int {
	var hasil int = 1
	var i int
	for i = 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Print("masukkan nilai a : ")
	fmt.Scan(&a)
	fmt.Print("masukkan nilai b : ")
	fmt.Scan(&b)
	fmt.Print("masukkan nilai c : ")
	fmt.Scan(&c)
	fmt.Print("masukkan nilai d : ")
	fmt.Scan(&d)

	fmt.Printf("hasil permutasi %v dan %v adalah : %v", a, c, permutasi(a, c))
	fmt.Println()
	fmt.Printf("hasil kombinasi %v dan %v adalah : %v", a, c, kombinasi(a, c))
	fmt.Println()

	fmt.Printf("hasil permutasi %v dan %v adalah : %v", b, d, permutasi(b, d))
	fmt.Println()
	fmt.Printf("hasil kombinasi %v dan %v adalah : %v", b, d, kombinasi(b, d))

}
