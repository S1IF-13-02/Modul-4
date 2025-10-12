package main

import "fmt"

func main() {
	var totalBelanja, diskon int

	fmt.Print("Masukkan total Belanja: ")
	fmt.Scanln(&totalBelanja)
	fmt.Print("Masukkan diskon: ")
	fmt.Scanln(&diskon)

	Harga := totalBelanja - (totalBelanja * diskon / 100)

	fmt.Printf("Total belanja akhir : %d\n", Harga)

}
