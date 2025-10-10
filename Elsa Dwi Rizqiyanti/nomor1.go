package main

import "fmt"

func main() {
	var totalBelanja, diskon int
	fmt.Scan(&totalBelanja)
	fmt.Scan(&diskon)

	totalAkhir := totalBelanja - (totalBelanja * diskon / 100)

	fmt.Println(totalAkhir)
}
