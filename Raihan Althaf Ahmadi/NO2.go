package main

import "fmt"
func main() {
	var BMI, berat, tinggi float64
	fmt.Print("Masukkan BMI: ")
	fmt.Scanln(&BMI)
	fmt.Print("Masukkan tinggi badan (m): ")
	fmt.Scanln(&tinggi)
	berat = BMI * (tinggi * tinggi)
	fmt.Printf("Indeks Massa Tubuh (BMI) Anda adalah: %.0f\n", berat)
}