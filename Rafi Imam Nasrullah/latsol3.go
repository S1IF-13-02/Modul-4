package main

import (
	"fmt"
)


func main() {
	var berat, tinggi, BMI float64

	fmt.Print("Masukan berat : ")
	fmt.Scanln(&berat)
	fmt.Print("Masukan tinggi : ")
	fmt.Scan(&tinggi)

	BMI = berat / (tinggi*tinggi)

	fmt.Printf("BMI = %.2f", BMI)

}