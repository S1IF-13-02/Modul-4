package main

import (
	"fmt"
	"math"
)

func hitungJarak(xa, ya, xb, yb float64) float64 {
	selisihXKuadrat := math.Pow(xb-xa, 2)
	selisihYKuadrat := math.Pow(yb-ya, 2)
	jarak := math.Sqrt(selisihXKuadrat + selisihYKuadrat)
	return jarak
}

func main() {
	var x1, y1, x2, y2, x3, y3 float64

	fmt.Scanln(&x1, &y1)
	fmt.Scanln(&x2, &y2)
	fmt.Scanln(&x3, &y3)

	sisiAB := hitungJarak(x1, y1, x2, y2)
	sisiBC := hitungJarak(x2, y2, x3, y3)
	sisiAC := hitungJarak(x1, y1, x3, y3)

	sisiTerpanjang := math.Max(sisiAB, math.Max(sisiBC, sisiAC))

	fmt.Printf("%.2f\n", sisiTerpanjang)
}