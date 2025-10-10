package main

import "fmt"

func main() {

	var berat, tinggi float64

	fmt.Print("masukan berat badan : ")
	fmt.Scan(&berat)
	fmt.Print("masukan tinggi badan : ")
	fmt.Scan(&tinggi)

	bmi := berat / (tinggi * tinggi)

	fmt.Printf("%.2f", bmi)
}
