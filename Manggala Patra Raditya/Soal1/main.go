package main

import "fmt"

func main() {
	var totalAwal, diskon int
	fmt.Scanln(&totalAwal)
	fmt.Scanln(&diskon)

	totalAkhir := totalAwal - (totalAwal * diskon / 100)
	fmt.Println(totalAkhir)
}
