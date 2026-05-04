package main

import "fmt"

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	if n%2 == 0 {
		// GENAP → horizontal
		for i := n; i >= 1; i-- {
			fmt.Printf("# %d ", i)
		}
		fmt.Println()
	} else {
		// GANJIL → vertikal
		for i := n; i >= 1; i-- {
			fmt.Printf("# %d %d -\n", i, i)
		}
	}
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
