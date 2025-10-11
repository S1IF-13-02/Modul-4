package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, x3, y3, AB, BC, CA, terpanjang float64

	fmt.Print("A: ")
	fmt.Scanln(&x1, &y1)
	fmt.Print("B: ")
	fmt.Scanln(&x2, &y2)
	fmt.Print("C: ")
	fmt.Scanln(&x3, &y3)
	AB = math.Sqrt(math.Pow((x2-x1), 2) + math.Pow((y2-y1), 2))
	BC = math.Sqrt(math.Pow((x3-x2), 2) + math.Pow((y3-y2), 2))
	CA = math.Sqrt(math.Pow((x1-x3), 2) + math.Pow((y1-y3), 2))
	fmt.Printf("panjang AB = %.2f\n", AB)
	fmt.Printf("panjang BC = %.2f\n", BC)
	fmt.Printf("panjang CA = %.2f\n", CA)
	if AB >= BC && AB >= CA {
		terpanjang = AB
	} else if BC > CA {
		terpanjang = BC
	} else {
		terpanjang = CA
	}
	fmt.Printf("sisi terpanjang = %.2f\n", terpanjang)
}
