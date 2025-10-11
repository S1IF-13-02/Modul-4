package main

import (
	"fmt"
	"math"
)

func main() {
	
	var bmi, tinggiBadan, beratBadan float64

	fmt.Print("Masukkan nilai BMI: ")
	fmt.Scan(&bmi)
	fmt.Print("Masukkan tinggi badan (m): ")
	fmt.Scan(&tinggiBadan)

	beratBadan = bmi * (tinggiBadan * tinggiBadan)
	fmt.Print("Berat badan anda adalah: ", math.Ceil(beratBadan), "Kg") 
}
