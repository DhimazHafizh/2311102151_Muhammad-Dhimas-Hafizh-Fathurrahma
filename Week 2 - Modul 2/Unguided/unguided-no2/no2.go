package main

import "fmt"

func main() {
	var p1, p2, p3, p4 string
	var i int = 1
	var berhasil bool = true

	for i <= 5 {
		fmt.Print("Percobaan ", i, " : ")
		fmt.Scan(&p1, &p2, &p3, &p4)
		if p1 != "merah" || p2 != "kuning" || p3 != "hijau" || p4 != "ungu" {
			berhasil = false
		}
		i++

	}
	fmt.Println("BERHASIL : ", berhasil)

}
