package main

import "fmt"

func main() {
	var input, d1, d2, d3 int

	fmt.Print("Masukkan 3 angka : ")
	fmt.Scan(&input)

	d1 = input / 100
	d2 = (input / 10) % 10
	d3 = input % 10

	fmt.Print(d1 <= d2 && d2 <= d3)

}
