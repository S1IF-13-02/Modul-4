package main

import "fmt"

func main() {
	var Berat, Tinggi float64
	fmt.Print("Masukkan Berat (kg): ")
	fmt.Scan(&Berat)
	fmt.Print("Masukkan Tinggi (m): ")
	fmt.Scan(&Tinggi)
	BMI := Berat / (Tinggi * Tinggi)
	fmt.Printf("BMI Anda: %.2f\n", BMI)
}