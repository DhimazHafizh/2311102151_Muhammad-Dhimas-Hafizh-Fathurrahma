package main

import "fmt"

func kuadrat(n int) int {
	return n * n
}

func fogoh(x int) int {
	return kuadrat((x + 1) - 2)
}

func gohof(x int) int {
	return (kuadrat(x) + 1) - 2
}

func hofog(x int) int {
	return kuadrat(x-2) + 1
}

func main() {
	var x, y, z int
	fmt.Print("Masukkan nilai x : ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y : ")
	fmt.Scan(&y)
	fmt.Print("Masukkan nilai z : ")
	fmt.Scan(&z)
	fmt.Println()

	fmt.Println("F(G(H(", x, "))) : ", fogoh(x))
	fmt.Println("G(H(F(", y, "))) : ", gohof(y))
	fmt.Println("H(F(G(", z, "))) : ", hofog(z))
}
