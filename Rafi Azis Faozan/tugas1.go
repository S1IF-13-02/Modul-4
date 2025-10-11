package main

import "fmt"

func main() {
	var HargaBarang float64
	var Diskon float64
	fmt.Print("Masukkan harga barang: ")
	fmt.Scan(&HargaBarang)
	fmt.Print("Masukkan diskon: ")
	fmt.Scan(&Diskon)
	PotonganHarga := HargaBarang * Diskon
	HargaAkhir := HargaBarang - PotonganHarga
	fmt.Print("Harga akhir: ", HargaAkhir)

}
