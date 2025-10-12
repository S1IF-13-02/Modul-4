package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, x3, y3 float64

	fmt.Println("Masukan Kordinat titik A : ")
	fmt.Scan(&x1, &y1)
	fmt.Println("Masukan Kordinat titik B : ")
	fmt.Scan(&x2, &y2)
	fmt.Println("Masukan Kordinat titik C : ")
	fmt.Scan(&x3, &y3)

	ab := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	bc := math.Sqrt(math.Pow(x3-x2, 2) + math.Pow(y3-y2, 2))
	ca := math.Sqrt(math.Pow(x3-x1, 2) + math.Pow(y3-y1, 2))

	terpanjang := ab
	if bc > terpanjang {
		terpanjang = bc
	}
	if ca > terpanjang {
		terpanjang = ca
	}

	fmt.Println("Sisi terpanjangnya Adalah: ")
	fmt.Printf("%.2f\n", terpanjang)
}