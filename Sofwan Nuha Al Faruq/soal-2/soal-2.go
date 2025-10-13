package main

import "fmt"

// Fungsi utama untuk menghitung berat badan berdasarkan BMI dan tinggi
func main() {
	// Variabel untuk menyimpan nilai BMI (Body Mass Index)
	var BMI float64
	// Variabel untuk menyimpan tinggi badan dalam meter
	var tinggi float64
	// Variabel untuk menyimpan hasil perhitungan berat badan dalam kg
	var berat float64

	// Meminta input nilai BMI dari pengguna
	fmt.Print("Masukkan nilai BMI: ")
	fmt.Scanln(&BMI)

	// Meminta input tinggi badan dalam meter dari pengguna
	fmt.Print("Masukkan tinggi badan (meter): ")
	fmt.Scanln(&tinggi)

	// Menghitung berat badan menggunakan rumus: Berat = BMI × (tinggi)²
	berat = BMI * (tinggi * tinggi)

	// Menampilkan hasil berat badan dengan format tanpa desimal
	fmt.Printf("Berat badan adalah: %.0f\n", berat)
}
