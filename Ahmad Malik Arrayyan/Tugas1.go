package main

import (
	"fmt"
)

func main() {
	var totalAwal int
	var diskon int

	// Input
	fmt.Scanln(&totalAwal)
	fmt.Scanln(&diskon)

	// Hitung total setelah diskon
	totalAkhir := totalAwal - (totalAwal * diskon / 100)

	// Output
	fmt.Println(totalAkhir)
}