package main 

import "fmt"

func main() {
	
	var hargaAwal, diskon, hargaAkhir int

	fmt.Print("Masukkan harga belanja awal: Rp")
	fmt.Scan(&hargaAwal)
	fmt.Print("Masukkan diskon (%): ")
	fmt.Scan(&diskon)

	hargaAkhir = hargaAwal * (100 - diskon) / 100
	fmt.Printf("Total harga belanja setelah diskon adalah: Rp%d", hargaAkhir)

}
