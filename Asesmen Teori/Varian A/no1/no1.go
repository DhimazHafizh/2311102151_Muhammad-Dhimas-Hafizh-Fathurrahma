package main

import "fmt"

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	for i := n; i >= 1; i-- {
		fmt.Println("#", i)
	}
}

/*

misal, jika input nya 10:
# 10
# 9
# 8
# 7
# 6
# 5
# 4
# 3
# 2
# 1

*/
