package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, x3, y3 float64
	fmt.Scan(&x1, &y1)
	fmt.Scan(&x2, &y2)
	fmt.Scan(&x3, &y3)

	sisiab := math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
	sisibc := math.Sqrt(math.Pow(x2-x3, 2) + math.Pow(y2-y3, 2))
	sisica := math.Sqrt(math.Pow(x3-x1, 2) + math.Pow(y3-y1, 2))

	semuasisi := (sisiab + sisibc + math.Sqrt(math.Pow(sisiab-sisibc, 2))) / 2
	panjangsisiterpanjang := (semuasisi + sisica + math.Sqrt(math.Pow(semuasisi-sisica, 2))) / 2

	fmt.Printf("%.2f\n", panjangsisiterpanjang)
}
