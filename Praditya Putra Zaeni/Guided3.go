package main

import "fmt"
 
func main() {
	var berat, tinggi float64
	fmt.Print("Masukkan berat: ")
	fmt.Scan(&berat)
	fmt.Print("Masukkan tinggi: ")
	fmt.Scan(&tinggi) 

	bmi := berat / (tinggi * tinggi)

	fmt.Printf("BMI: %.2f\n", bmi)
}