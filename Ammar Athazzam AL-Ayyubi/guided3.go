package main

import (
	"fmt"
)

func main() {
	var berat, tinggi, bmi float64
	fmt.Print("berat dan tinggi: ")
	fmt.Scan(&berat)
	fmt.Scan(&tinggi)

	bmi = berat / (tinggi * tinggi)
	fmt.Printf("%.2f", bmi)
}
