package main

import "fmt"

func main() {
	var bmi, tinggi, berat float64
	fmt.Print("Masukan bmi :")
	fmt.Scan(&bmi)
	fmt.Print("Masukan tinggi :")
	fmt.Scan(&tinggi)

	berat = bmi * tinggi * tinggi

	fmt.Printf("Berat badan kamu: %.0f kg\n", berat)
}
