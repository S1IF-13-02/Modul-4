package main

import (
	"fmt"
	"math"
)

func HitungBeratBadan(bmi, tb float64) float64 {
	return bmi * math.Pow(tb, 2)
}

func main() {
	var bmi, tb float64

	fmt.Print("Masukkan BMI: ")
	fmt.Scan(&bmi)

	fmt.Print("Masukkan Tinggi Badan: ")
	fmt.Scan(&tb)

	result := HitungBeratBadan(bmi, tb)

	fmt.Println(math.Ceil(result))
}
