package main

import (
	"fmt"
)

func main() {
	var bmi, tinggi, berat float64

	// Input nilai BMI dan tinggi badan
	fmt.Scanln(&bmi, &tinggi)

	// Hitung berat badan
	berat = bmi * tinggi * tinggi

	// Output hasil
	fmt.Printf("%.0f\n", berat)
}