package main

import (
	"fmt"
	"math"
)

func hitungJarak(x1, y1, x2, y2 float64) float64 {
	deltaXKuadrat := math.Pow(x2-x1, 2)
	deltaYKuadrat := math.Pow(y2-y1, 2)
	return math.Sqrt(deltaXKuadrat + deltaYKuadrat)
}

func main() {

	var ax, ay, bx, by, cx, cy float64

	fmt.Scan(&ax, &ay)

	fmt.Scan(&bx, &by)

	fmt.Scan(&cx, &cy)

	sisiAB := hitungJarak(ax, ay, bx, by)
	sisiBC := hitungJarak(bx, by, cx, cy)
	sisiCA := hitungJarak(cx, cy, ax, ay)

	sisiTerpanjang := sisiAB

	if sisiBC > sisiTerpanjang {
		sisiTerpanjang = sisiBC
	}

	if sisiCA > sisiTerpanjang {
		sisiTerpanjang = sisiCA
	}
	fmt.Printf("%.2f\n", sisiTerpanjang)
}
