package main

import "fmt"

func main() {
	var bmi, tinggi_badan float64
	fmt.Print("Nilai BMI: ")
	fmt.Scan(&bmi)
	fmt.Print("Tinggi badan: ")
	fmt.Scan(&tinggi_badan)

	berat_badan := bmi * (tinggi_badan * tinggi_badan)

	fmt.Printf("Berat Badan: %.f", berat_badan)
}
