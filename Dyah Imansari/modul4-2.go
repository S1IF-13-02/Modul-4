package main

import "fmt"

func main() {
	var BMI, tinggi, berat float64

	fmt.Print("Masukkan BMI:")
	fmt.Scanln(&BMI)
	fmt.Print("Masukkan tinggi:")
	fmt.Scanln(&tinggi)
	berat = BMI * tinggi * tinggi
	fmt.Printf("berat badan: %.0fkg\n", berat)

}
