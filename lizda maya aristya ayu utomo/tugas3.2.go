package main

import "fmt"

func main() {
	var bmi, tinggi float64
	fmt.Print("Masukkan nilai bmi: ")
	fmt.Scan(&bmi)
	fmt.Print("Masukkan tinggi: ")
	fmt.Scan(&tinggi)

	berat := bmi * (tinggi * tinggi)

	fmt.Printf("berat: %.0f kg\n", berat)
}
