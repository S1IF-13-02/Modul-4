package main

import (
	"math"

	"fmt"
)

func main() {

	var ax, ay, bx, by, cx, cy float64

	fmt.Print("masukan koordinat ax , ay : ")
	fmt.Scan(&ax, &ay)
	fmt.Print("masukan koordunat bx , by : ")
	fmt.Scan(&bx, &by)
	fmt.Print("masukan koordinat cx , cy : ")
	fmt.Scan(&cx, &cy)

	ab := math.Sqrt(math.Pow(bx-ax, 2) + math.Pow(by-ay, 2))
	bc := math.Sqrt(math.Pow(cx-bx, 2) + math.Pow(cy-by, 2))
	ca := math.Sqrt(math.Pow(ax-cx, 2) + math.Pow(ay-cy, 2))

	sisiterpanjang := math.Max(ab, math.Max(bc, ca))

	fmt.Printf("sisi terpanjang: %.2f", sisiterpanjang)

}
