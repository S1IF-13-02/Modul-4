package main

import (
	"fmt"
	"math"
)

func main() {

	var ax, ay, bx, by, cx, cy float64

	fmt.Print("Masukkan koordinat (ax ay): ")
	fmt.Scan(&ax, &ay)
	fmt.Print("Masukkan koordinat (bx by): ")
	fmt.Scan(&bx, &by)
	fmt.Print("Masukkan koordinat (cx cy): ")
	fmt.Scan(&cx, &cy)

	ab := math.Sqrt(math.Pow(bx-ax, 2) + math.Pow(by-ay, 2))
	bc := math.Sqrt(math.Pow(cx-bx, 2) + math.Pow(cy-by, 2))
	ca := math.Sqrt(math.Pow(ax-cx, 2) + math.Pow(ay-cy, 2))

	var terpanjang float64 = ab

	if bc > terpanjang {
		terpanjang = bc
	}
	if ca > terpanjang {
		terpanjang = ca
	}

	fmt.Printf("Sisi AB: %.2f\n", ab)
	fmt.Printf("Sisi BC: %.2f\n", bc)
	fmt.Printf("Sisi CA: %.2f\n", ca)
	fmt.Printf("Sisi terpanjang adalah: %.2f\n", terpanjang)
}
