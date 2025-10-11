package main

import "fmt"

func main() {
	var totalBelanja, diskon, totalAkhir int

	fmt.Scan(&totalBelanja)
	fmt.Scan(&diskon)
	totalAkhir = totalBelanja - (totalBelanja * diskon / 100)
	fmt.Println(totalAkhir)

	fmt.Scan(&totalBelanja)
	fmt.Scan(&diskon)
	totalAkhir = totalBelanja - (totalBelanja * diskon / 100)
	fmt.Println(totalAkhir)

	fmt.Scan(&totalBelanja)
	fmt.Scan(&diskon)
	totalAkhir = totalBelanja - (totalBelanja * diskon / 100)
	fmt.Println(totalAkhir)
}
