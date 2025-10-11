package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, x3, y3 float64
	fmt.Print("Masukkan x1 dan y1: ")
	fmt.Scan(&x1, &y1)
	fmt.Print("Masukkan x2 dan y2: ")
	fmt.Scan(&x2, &y2)
	fmt.Print("Masukkan x3 dan y3: ")
	fmt.Scan(&x3, &y3)
	AB := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	BC := math.Sqrt(math.Pow(x3-x2, 2) + math.Pow(y3-y2, 2))
	CA := math.Sqrt(math.Pow(x1-x3, 2) + math.Pow(y1-y3, 2))
	sisi_terpanjang := AB
	if BC > sisi_terpanjang {
		sisi_terpanjang = BC
	}
	if CA > sisi_terpanjang {
		sisi_terpanjang = CA
	}
	fmt.Printf("Sisi terpanjang pada segitiga siku siku: %.2f\n", sisi_terpanjang)
}
