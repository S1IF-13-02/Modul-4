package main

import "fmt"

func main() {
	var totalBelanja float64
	var diskonPersen float64

	fmt.Print("Masukkan totalBelanja: ")
	fmt.Scan(&totalBelanja)
	fmt.Print("Masukkan diskonPersen: ")
	fmt.Scan(&diskonPersen)

	potonganHarga := totalBelanja * (diskonPersen / 100)
	hargaAkhir := totalBelanja - potonganHarga
	fmt.Println(int(hargaAkhir))
}
