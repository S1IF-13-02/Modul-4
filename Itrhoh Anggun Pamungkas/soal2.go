package main

import "fmt"

func main() {
	var bmi, tinggi, berat float64

	fmt.Print("BMI: ")
	fmt.Scan(&bmi)
	fmt.Print("tinggi: " )
	fmt.Scan(&tinggi)
	berat = bmi * tinggi * tinggi

	fmt.Printf("Berat badan: %.0f kg\n", berat)
}
