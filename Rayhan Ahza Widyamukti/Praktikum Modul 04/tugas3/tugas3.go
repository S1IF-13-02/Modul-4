package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, x3, y3 float64

	fmt.Println("Masukan Kordinat titik A, B, dan C: ")
	fmt.Scan(&x1, &y1)
	fmt.Scan(&x2, &y2)
	fmt.Scan(&x3, &y3)

	fmt.Println("Masukan Sisi-Sisi Segitiga")
	ab := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	bc := math.Sqrt(math.Pow(x3-x2, 2) + math.Pow(y3-y2, 2))
	ca := math.Sqrt(math.Pow(x3-x1, 2) + math.Pow(y3-y1, 2))

	// Tentukan sisi terpanjang
	terpanjang := ab
	if bc > terpanjang {
		terpanjang = bc
	}
	if ca > terpanjang {
		terpanjang = ca
	}

	fmt.Println("Hasilnya Adalah: ")
	fmt.Printf("%.2f\n", terpanjang)
}
