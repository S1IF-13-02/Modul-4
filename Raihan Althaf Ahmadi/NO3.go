package main

import (
	"fmt"
	"math"
)

func dist(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func main() {
	var ax, ay, bx, by, cx, cy float64
	fmt.Print("Masukkan koordinat titik A (x y): ")
	fmt.Scanf("%f %f\n", &ax, &ay)
	fmt.Print("Masukkan koordinat titik B (x y): ")
	fmt.Scanf("%f %f\n", &bx, &by)
	fmt.Print("Masukkan koordinat titik C (x y): ")
	fmt.Scanf("%f %f\n", &cx, &cy)

	ab := dist(ax, ay, bx, by)
	bc := dist(bx, by, cx, cy)
	ca := dist(cx, cy, ax, ay)

	terpanjang := math.Max(ab, math.Max(bc, ca))

	fmt.Printf("%.2f\n", terpanjang)
}
