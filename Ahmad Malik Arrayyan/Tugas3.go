package main

import (
	"fmt"
	"math"
)

func main() {
	var ax, ay, bx, by, cx, cy float64

	// Input koordinat titik A, B, dan C
	fmt.Scanln(&ax, &ay)
	fmt.Scanln(&bx, &by)
	fmt.Scanln(&cx, &cy)

	// Hitung panjang sisi-sisi segitiga
	ab := math.Sqrt(math.Pow(bx-ax, 2) + math.Pow(by-ay, 2))
	bc := math.Sqrt(math.Pow(cx-bx, 2) + math.Pow(cy-by, 2))
	ca := math.Sqrt(math.Pow(ax-cx, 2) + math.Pow(ay-cy, 2))

	// Tentukan sisi terpanjang
	maxSide := ab
	if bc > maxSide {
		maxSide = bc
	}
	if ca > maxSide {
		maxSide = ca
	}

	// Tampilkan hasil dengan dua angka di belakang koma
	fmt.Printf("%.2f\n", maxSide)
}