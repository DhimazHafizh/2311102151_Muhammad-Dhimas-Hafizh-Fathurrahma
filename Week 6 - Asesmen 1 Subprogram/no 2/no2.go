package main

import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	/* Mengembalikan jumlah hari maksimum yang biaya perjalanannya ditanggung
	   oleh Tel-U berdasarkan lama hari study tour (jumlahHari) dan tujuan (domestik/mancanegara) */
	if tujuan == "domestik" && jumlahHari > 3 {
		jumlahHari = 3 //maks hari domestik yang biayanya ditanggung = 3 hari
	} else if tujuan == "mancanegara" && jumlahHari > 8 {
		jumlahHari = 8 //maks hari mancanegara yang biayanya ditanggung = 8 hari
	}
	return jumlahHari
}

func biayaPerHari(jumlahMhs int) int {
	/* Menghitung biaya tour domestik per hari yang ditanggung oleh Tel-U untuk
	   jumlah mahasiswa sebanyak jumlahMhs */

	//yang dihitung biayaPerHari hanya untuk domestik saja, karena untuk mancanegara biayaPerHari nya = biayaPerHari domestik x 1,5 (1,5 kali biaya domestik)

	//biaya makan siang + makan malam = 35.000 + 35.000 = 70.000
	//biaya penginapan = 250.000
	//uang saku = 300.000
	return jumlahMhs * (70000 + 250000 + 300000)

}

func perhitunganBiaya(jumlahMhs, lamaPerjalanan int, tujuan string, totalBiaya *float64) { //parameter totalBiaya menggunakan pass by reference
	/* I.S. Terdefinisi jumlah mahasiswa, lama hari study tour, dan tujuan perjalanan (domestik/mancanegara)
	   F.S. Telah dihitung biaya perjalanan yang ditanggung Tel-U */

	// Panggil salah satu fungsi/prosedur untuk menghitung lama perjalanan
	var lama int = tanggunganHari(lamaPerjalanan, tujuan)

	// Panggil fungsi/prosedur untuk menghitung biaya total tour domestik seluruh mahasiswa
	*totalBiaya = float64(biayaPerHari(jumlahMhs))

	// Hitung biaya study tour seluruh mahasiswa jika tujuan domestik atau mancanegara
	*totalBiaya = *totalBiaya * float64(lama)
	if tujuan == "mancanegara" {
		*totalBiaya = 1.5 * *totalBiaya
	}
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64 = 0

	// lakukan proses masukan atau input di sini
	fmt.Print("masukkan jumlah mahasiswa : ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukkan lama hari study tour : ")
	fmt.Scan(&lama)
	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara) : ")
	fmt.Scan(&tujuan)
	fmt.Println()

	// hitung biaya perjalanan yang dikeluarkan Tel-U dengan memanggil subprogram yang tepat
	perhitunganBiaya(jumlah, lama, tujuan, &biaya) //biaya memakai pass by reference

	// tampilkan biaya
	fmt.Printf("Biaya perjalanan yang harus dikeluarkan Tel-U : Rp. %.0f", biaya)
}
