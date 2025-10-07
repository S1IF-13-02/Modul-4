package main

import (
	"fmt"
)

func main() {
	var totalAwal int
	var diskon int

	fmt.Print("Masukkan total belanja awal: ")
	fmt.Scanln(&totalAwal)

	fmt.Print("Masukkan besar diskon (%): ")
	fmt.Scanln(&diskon)

	totalAkhir := totalAwal - (totalAwal * diskon / 100)

	fmt.Println("Total belanja setelah diskon adalah:", totalAkhir)

	
}
