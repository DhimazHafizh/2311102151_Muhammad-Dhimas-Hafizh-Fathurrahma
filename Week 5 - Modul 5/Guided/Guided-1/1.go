package main

import "fmt"

func baris(bilangan int) {
	if bilangan == 1 {
		fmt.Println(1)
	} else {
		fmt.Println(bilangan)
		baris(bilangan - 1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan input : ")
	fmt.Scan(&n)
	baris(n)
}

//input = 5
//5 (bilangan sekarang = 5)
//(bilangan - 1) 5 - 1 = 4 (bilangan sekarang = 4)
//(bilangan - 1) 4 - 1 = 3 (bilangan sekarang = 3)
//(bilangan - 1) 3 - 1 = 2 (bilangan sekarang = 2)
//(bilangan - 1) 2 - 1 = 1 (bilangan sekarang = 1)
