package main

import (
	"fmt"
)

func main() {
	var berat, tinggi float64
	fmt.Print("input Berat  (kg): ")
	fmt.Scan(&berat)

	fmt.Print("input Tinggi  (cm): ")
	fmt.Scan(&tinggi)

	mtinggi := tinggi / 100

	bmi := berat / (mtinggi * mtinggi)
	fmt.Printf("BMI : %.2f", bmi)
}
