package main

import (
	"fmt"
	"math"
)

func main() {

	var ax, ay float64
	var bx, by float64
	var cx, cy float64

	fmt.Scanln(&ax, &ay)
	fmt.Scanln(&bx, &by)
	fmt.Scanln(&cx, &cy)

	ab := math.Sqrt(math.Pow(ax-bx, 2) + math.Pow(ay-by, 2))
	bc := math.Sqrt(math.Pow(bx-cx, 2) + math.Pow(by-cy, 2))
	ca := math.Sqrt(math.Pow(cx-ax, 2) + math.Pow(cy-ay, 2))

	palingPanjang := ab
	if bc > palingPanjang {
		palingPanjang = bc
	}
	if ca > palingPanjang {
		palingPanjang = ca
	}

	fmt.Println("\n--- hasil ---")
	fmt.Printf("hasil: %.2f\n", palingPanjang)
}
