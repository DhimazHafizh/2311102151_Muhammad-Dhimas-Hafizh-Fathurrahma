package main

import "fmt"

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)

	tahun := 2012

	for i := 0; i < n; i++ {
		fmt.Println(tahun + i)
	}
}

/*
misal, input 5, maka outputnya:
2012
2013
2014
2015
2016

misal, jika input nya 9:
2012
2013
2014
2015
2016
2017
2018
2019
2020

*/
