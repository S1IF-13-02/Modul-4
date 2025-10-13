package main

import (
	"fmt"
	"math"
)

func main() {
	var titikAx, titikAy, titikBx, titikBy, titikCx, titikCy float64

	// Input koordinat titik A, B, dan C
	fmt.Print("Masukkan koordinat titik A (x y): ")
	fmt.Scanln(&titikAx, &titikAy)

	fmt.Print("Masukkan koordinat titik B (x y): ")
	fmt.Scanln(&titikBx, &titikBy)

	fmt.Print("Masukkan koordinat titik C (x y): ")
	fmt.Scanln(&titikCx, &titikCy)

	// Hitung panjang sisi-sisi segitiga
	sisiAB := math.Sqrt(math.Pow(titikBx-titikAx, 2) + math.Pow(titikBy-titikAy, 2))
	sisiBC := math.Sqrt(math.Pow(titikCx-titikBx, 2) + math.Pow(titikCy-titikBy, 2))
	sisiCA := math.Sqrt(math.Pow(titikAx-titikCx, 2) + math.Pow(titikAy-titikCy, 2))

	// Tentukan sisi terpanjang
	sisiTerpanjang := sisiAB
	if sisiBC > sisiTerpanjang {
		sisiTerpanjang = sisiBC
	}
	if sisiCA > sisiTerpanjang {
		sisiTerpanjang = sisiCA
	}

	// Tampilkan hasil dengan dua angka di belakang koma
	fmt.Printf("Sisi terpanjang: %.2f\n", sisiTerpanjang)
}
