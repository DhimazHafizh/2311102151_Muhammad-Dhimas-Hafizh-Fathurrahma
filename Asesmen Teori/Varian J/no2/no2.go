package main

import "fmt"

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	tahun := 2012

	for i := 0; i < n; i++ {
		t := tahun + i

		if t%4 == 0 {
			fmt.Println(t, "merupakan tahun kabisat")
		} else {
			fmt.Println(t)
		}
	}
}

/*
misal, input 5, outputnya:
2012 merupakan tahun kabisat
2013
2014
2015
2016 merupakan tahun kabisat

misal, input 9, outputnya:
2012 merupakan tahun kabisat
2013
2014
2015
2016 merupakan tahun kabisat
2017
2018
2019
2020 merupakan tahun kabisat
*/
