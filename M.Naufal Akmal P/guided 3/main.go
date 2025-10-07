package main

import "fmt"

func main () {
	var  berat, tinggi float64
	
	fmt.Print("masukan berat (kg): ")
	fmt.Scan(&berat)
	fmt.Print("masukan tinggi (m): ")
	fmt.Scan(&tinggi)

	bmi := berat / (tinggi * tinggi)

	fmt.Printf("bmi: %.2f\n", bmi)
}