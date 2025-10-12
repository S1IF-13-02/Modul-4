package main

import (
	"fmt"
)

func main() {
	var berat, tinggi float64

	fmt.Print("Masukkan berat (kg): ")
	fmt.Scan(&berat)
	fmt.Print("Masukkan tinggi (m): ")
	fmt.Scan(&tinggi)

	BMI := berat / (tinggi * tinggi)

	fmt.Printf("Hasil BMI Anda adalah: %.2f\n", BMI)
}
