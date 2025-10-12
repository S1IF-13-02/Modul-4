package main

import (
	"fmt"
	"math"
)

func main() {
	var BMI float64
	var tinggi float64

	fmt.Print("Masukkan BMI: ")
	fmt.Scan(&BMI)
	fmt.Print("Masukkan tinggi: ")
	fmt.Scan(&tinggi)

	var berat = BMI * (tinggi * tinggi)
	fmt.Println(math.Round(berat))

}
