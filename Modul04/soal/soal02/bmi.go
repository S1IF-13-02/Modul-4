package main

import (
	"fmt"
	"math"
)

func main() {
	var bmi, tinggiBadan float64

	fmt.Print("Masukkan nilai BMI: ")
	fmt.Scan(&bmi)
	fmt.Print("Masukkan tinggi badan (meter): ")
	fmt.Scan(&tinggiBadan)

	beratBadan := bmi * (tinggiBadan * tinggiBadan)
	beratBadanBulat := math.Round(beratBadan)
	fmt.Printf("Berat badan anda adalah: %.0f kg\n", beratBadanBulat)
}
