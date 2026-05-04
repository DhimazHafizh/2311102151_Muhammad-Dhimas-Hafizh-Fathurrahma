package main

import "fmt"

func main() {
	cetak(5)
}
func cetak(x int) {
	if x == 10 {
		fmt.Println(x)
	} else {
		cetak(x + 1)
		fmt.Println(x)
	}
}
