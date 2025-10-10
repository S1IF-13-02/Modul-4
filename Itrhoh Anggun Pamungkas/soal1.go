package main

import "fmt"

func main() {
	var total, diskon int

	fmt.Print("harga: ")
	fmt.Scan(&total)
	fmt.Print("diskon: ")
	fmt.Scan(&diskon)

	totalAkhir := total - (total * diskon / 100)

	fmt.Println("total: ", totalAkhir)
}