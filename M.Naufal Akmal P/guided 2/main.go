package main

import "fmt"

func main () {
	var angka int
	fmt.Print("masukan angka:")
	fmt.Scan(&angka)

	d1 := angka / 100
	d2 := (angka / 10) %10
	d3 := angka % 10

	hasilnya:= d1 < d2 && d2 < d3
	fmt.Println(hasilnya)

}