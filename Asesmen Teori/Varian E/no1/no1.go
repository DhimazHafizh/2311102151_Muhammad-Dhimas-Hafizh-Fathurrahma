package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Println(i, ". *")
	}
}

/*
misal, jika input nya 5,outputnya:
1. *
2. *
3. *
4. *
5. *

misal, jika input nya 10, outputnya:
1. *
2. *
3. *
4. *
5. *
6. *
7. *
8. *
9. *
10. *

*/
