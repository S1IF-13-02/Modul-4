package main

import "fmt"

func HitungDiskon(total_harga_awal, diskon int) int {
	harga_diskon := (total_harga_awal * diskon) / 100
	return total_harga_awal - harga_diskon
}

func main() {
	var total_harga_awal, diskon int

	fmt.Print("Masukkan total harga barang: ")
	fmt.Scan(&total_harga_awal)

	fmt.Print("Masukkan diskon: ")
	fmt.Scan(&diskon)

	result := HitungDiskon(total_harga_awal, diskon)

	fmt.Println("Total harga akhir: ", result)
}
