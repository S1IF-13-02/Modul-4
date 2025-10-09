package main

import "fmt"

func main() {

	var (
		harga, diskon int
	)

	fmt.Print("Masukan harga belanja : ")
	fmt.Scan(&harga)
	fmt.Print("masukan diskon : ")
	fmt.Scan(&diskon)

	totalakhir := harga - (harga * diskon / 100)

	fmt.Println("Total setelah diskon:", totalakhir)

}
