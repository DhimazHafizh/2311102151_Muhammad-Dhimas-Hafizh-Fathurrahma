package main

import "fmt"

func hitungPersegi(sisi int) {
	fmt.Println("Luas persegi : ", sisi*sisi)
	fmt.Println("Keliling persegi : ", 4*sisi)
}

func hitungPersegiPanjang(panjang, lebar int) {
	fmt.Println("Luas persegi panjang : ", panjang*lebar)
	fmt.Println("Keliling persegi panjang : ", 2*(panjang+lebar))
}

func hitungLingkaran(jarijari float64) {
	fmt.Println("Luas lingkaran : ", 3.14*(jarijari*jarijari))
	fmt.Println("Keliling lingkaran : ", 2*3.14*jarijari)
}

func main() {
	var pilihan, s, p, l int
	var jari float64

	fmt.Println("--- PROGRAM BANGUN DATAR ---")
	fmt.Println("1. Hitung luas & keliling persegi")
	fmt.Println("2. Hitung luas & keliling persegi panjang")
	fmt.Println("3. Hitung luas & keliling lingkaran")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilihan)
	fmt.Println()

	switch pilihan {
	case 1:
		fmt.Print("Masukkan sisi : ")
		fmt.Scan(&s)
		hitungPersegi(s)
	case 2:
		fmt.Print("Masukkan panjang : ")
		fmt.Scan(&p)
		fmt.Print("Masukkan lebar : ")
		fmt.Scan(&l)
		hitungPersegiPanjang(p, l)
	case 3:
		fmt.Print("Masukkan jari-jari : ")
		fmt.Scan(&jari)
		hitungLingkaran(jari)
	default:
		fmt.Println("Pilihan tidak valid!")
	}
}
