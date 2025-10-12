package main

import (
	"fmt"
)

func main() {
	var bmi, tinggi, berat float64

	fmt.Scan(&bmi, &tinggi)
	berat = bmi * (tinggi * tinggi)
	fmt.Printf("%.0f\n", berat)
}
