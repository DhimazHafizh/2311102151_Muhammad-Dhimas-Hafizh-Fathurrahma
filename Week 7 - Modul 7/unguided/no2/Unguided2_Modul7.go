package main

import "fmt"

type angka int
type kata string

type Buku struct {
	judul         kata
	penulis       kata
	penerbit      kata
	tahunTerbit   angka
	jumlahHalaman angka
}

func main() {
	var biodataBuku Buku

	fmt.Println("=== INPUT BIODATA BUKU ===")
	fmt.Print("Masukkan judul buku : ")
	fmt.Scan(&biodataBuku.judul)
	fmt.Print("Masukkan nama penulis : ")
	fmt.Scan(&biodataBuku.penulis)
	fmt.Print("Masukkan penerbit : ")
	fmt.Scan(&biodataBuku.penerbit)
	fmt.Print("Masukkan tahun terbit : ")
	fmt.Scan(&biodataBuku.tahunTerbit)
	fmt.Print("Masukkan jumlah halaman: ")
	fmt.Scan(&biodataBuku.jumlahHalaman)
	fmt.Println()

	fmt.Println("=== BIODATA BUKU ===")
	fmt.Println("Judul Buku :", biodataBuku.judul)
	fmt.Println("Penulis :", biodataBuku.penulis)
	fmt.Println("Penerbit :", biodataBuku.penerbit)
	fmt.Println("Tahun Terbit :", biodataBuku.tahunTerbit)
	fmt.Println("Jumlah Halaman :", biodataBuku.jumlahHalaman)

}
