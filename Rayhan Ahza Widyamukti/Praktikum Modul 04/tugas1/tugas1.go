package main

import "fmt"

func main() {
	var totalBelanja, diskon int
	fmt.Println("Masukan Total Belanja: ")
	fmt.Scanln(&totalBelanja)
	fmt.Println("Masukan Diskon: ")
	fmt.Scanln(&diskon)

	totalAkhir := totalBelanja - (totalBelanja * diskon / 100)
	fmt.Println("Total Akhir: ")
	fmt.Println(totalAkhir)
}
