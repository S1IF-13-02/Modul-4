package main

import (
	"fmt"
	"math"
)

func main() {
	var bmi, tinggi float64
	fmt.Scan(&bmi)
	fmt.Scan(&tinggi)

	berat := bmi * (tinggi * tinggi)
	beratBulat := math.Round(berat)

	fmt.Printf("Berat badan: %.0f kg\n", beratBulat)
}
