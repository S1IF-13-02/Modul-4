package main

import (
	"fmt"
	"math"
)

func hitungJarak(x1, y1, x2, y2 float64) float64 {
	jarakKuadrat := math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2)
	return math.Sqrt(jarakKuadrat)
}

func main() {
	var x1, y1, x2, y2, x3, y3 float64

	fmt.Println("Masukkan koordinat Titik A (x y):")
	fmt.Scan(&x1, &y1)
	fmt.Println("Masukkan koordinat Titik B (x y):")
	fmt.Scan(&x2, &y2)
	fmt.Println("Masukkan koordinat Titik C (x y):")
	fmt.Scan(&x3, &y3)

	sisiAB := hitungJarak(x1, y1, x2, y2)
	sisiBC := hitungJarak(x2, y2, x3, y3)
	sisiCA := hitungJarak(x3, y3, x1, y1)

	sisiTerpanjangSementara := math.Max(sisiAB, sisiBC)
	sisiTerpanjang := math.Max(sisiTerpanjangSementara, sisiCA)

	fmt.Printf("\nPanjang sisi terpanjang adalah: %.2f\n", sisiTerpanjang)
}
