package main

import "fmt"

// function sesuai perintah soal
func cetakAngka(n int) {
	if n%2 == 0 {
		// genap → horizontal
		for i := n; i >= 1; i-- {
			fmt.Printf("# %d ", i)
		}
		fmt.Println()
	} else {
		// ganjil → vertikal
		for i := n; i >= 1; i-- {
			fmt.Printf("# %d %d -\n", i, i)
		}
	}
}

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	// panggil function
	cetakAngka(n)
}

/*
misal, input 10 (genap):
# 10 # 9 # 8 # 7 # 6 # 5 # 4 # 3 # 2 # 1

misal, input 5 (ganjil):
# 5 5 -
# 4 4 -
# 3 3 -
# 2 2 -
# 1 1 -
*/
