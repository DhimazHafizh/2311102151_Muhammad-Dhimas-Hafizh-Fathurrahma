package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Print("#- ")
	}
}

/* 
misal, jika input nya 5:
#- #- #- #- #-

misal, jika input nya 8:
#- #- #- #- #- #- #- #-

*/