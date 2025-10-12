package main

import "fmt"

func main() {
	var totalBelanja, diskon int
	fmt.Print("Masukkan total belanja: ")
	fmt.Scanln(&totalBelanja)
	fmt.Print("Masukkan besaran diskon (dalam persen): ")
	fmt.Scanln(&diskon)
	totalSetelahDiskon := totalBelanja * (100 - diskon) / 100
	fmt.Println("Total belanja setelah diskon:", totalSetelahDiskon)
}
