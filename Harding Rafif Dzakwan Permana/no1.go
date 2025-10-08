package main

import "fmt"

func main () {
	var belanja, diskon float64

	fmt.Print("Masukan belanja total :")
	fmt.Scan(&belanja)

	fmt.Print("Masukan Diskon :")
	fmt.Scan(&diskon)

	Hasilakhir := belanja - (belanja* diskon / 100)

	fmt.Println(Hasilakhir)
}