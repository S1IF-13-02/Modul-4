package main

import "fmt"

func main() {
	var bmi, tinggi, berat float64
	fmt.Print("Masukan nilai BMI:")
	fmt.Scan(&bmi)
	fmt.Print("Masukan tinggi badan:")
	fmt.Scan(&tinggi)

	berat = bmi * (tinggi * tinggi)

	fmt.Printf("Berat badan anda adalah: %.0f kg", berat)
}