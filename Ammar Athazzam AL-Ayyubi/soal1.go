package main

import (
	"fmt"
)

func main() {
	var hargatotal int
	var diskon int

	fmt.Println("harga total: ")
	fmt.Scanln(&hargatotal)

	fmt.Printf("diskon: ")
	fmt.Scanln(&diskon)

	hargaakhir := hargatotal - (hargatotal * diskon / 100)

	fmt.Println("\n--- HASIL ---")
	fmt.Println("hargatotal:", hargatotal)
	fmt.Println("diskon:", diskon, "%")
	fmt.Println("setelah diskon:", hargaakhir)
}
