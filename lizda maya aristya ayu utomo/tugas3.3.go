package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, x3, y3 float64
	fmt.Print("masukkan kooridnat titik A:")
	fmt.Scan(&x1, &y1)
	fmt.Print("masukkan kooridnat titik B:")
	fmt.Scan(&x2, &y2)
	fmt.Print("masukkan kooridnat titik C:")
	fmt.Scan(&x3, &y3)

	SisiAB := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	SisiBC := math.Sqrt(math.Pow(x3-x2, 2) + math.Pow(y3-y2, 2))
	SisiCA := math.Sqrt(math.Pow(x1-x3, 2) + math.Pow(y1-y3, 2))

	terpanjang := SisiAB
	if SisiBC > terpanjang {
		terpanjang = SisiBC
	}
	if SisiCA > terpanjang {
		terpanjang = SisiCA
	}

	fmt.Printf("Sisi terpanjang: %.2f\n", terpanjang)
}
