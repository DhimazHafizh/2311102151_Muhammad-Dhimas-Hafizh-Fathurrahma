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

	// Hitung kg dan sisa gram
	kg = beratGram / 1000
	sisaGramkg = beratGram % 1000
	sisaGram = sisaGramkg

	// Biaya per kg
	biayakg = kg * 10000

	// Jika total berat lebih dari 10 kg, sisa gram gratis
	if kg > 10 {
		sisaGram = 0
	} else {
		// Hitung biaya sisa gram
		if sisaGram >= 500 {
			biayagram = sisaGram * 5
			// totalBiaya += sisaGram * 5
		} else {
			biayagram = sisaGram * 15
			// totalBiaya += sisaGram * 15
		}
	}
	totalBiaya = biayakg + biayagram

	// Output hasil
	fmt.Println("\n===== Detail Perhitungan =====")
	fmt.Println("Detail berat : ", kg, " kg + ", sisaGramkg, " gram")
	fmt.Println("Detail biaya  : Rp.", biayakg, " + Rp.", biayagram)
	fmt.Println("Total biaya: Rp", totalBiaya)
}
