package main

import (
	"fmt"
)

func main() {
	var beratGram int
	var kg, sisaGramkg, sisaGram int
	var biayakg, biayagram, totalBiaya int

	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scan(&beratGram)

	kg = beratGram / 1000
	sisaGramkg = beratGram % 1000
	sisaGram = sisaGramkg

	biayakg = kg * 10000

	if kg > 10 {
		sisaGram = 0
	} else {
		if sisaGram >= 500 {
			biayagram = sisaGram * 5
		} else {
			biayagram = sisaGram * 15
		}
	}
	totalBiaya = biayakg + biayagram

	fmt.Println()
	fmt.Println("===== Detail Perhitungan =====")
	fmt.Println("Detail berat : ", kg, " kg + ", sisaGramkg, " gram")
	fmt.Println("Detail biaya  : Rp.", biayakg, " + Rp.", biayagram)
	fmt.Println("Total biaya: Rp", totalBiaya)
}
