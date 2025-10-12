package main

import "fmt"

func main() {
	var bmi, tinggi, berat float64

	fmt.Println("Masukan nilai BMI dan Tinggi Badan: ")
	fmt.Scan(&bmi, &tinggi)

	berat = bmi * (tinggi * tinggi)

	fmt.Println("Hasil Berat Badan: ")
	fmt.Printf("%.0f\n", berat)
}
