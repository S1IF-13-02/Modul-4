package main

import "fmt"

func main() {
	var angka int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)
	d1 := angka / 1000
	d2 := (angka % 1000) / 100
	d3 := (angka % 100) / 10
	hasilnya := d1 < d2 && d2 < d3
	fmt.Println(hasilnya)
}