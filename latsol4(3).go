package main

import "fmt"

func main() {
	var berat, tinggi float64

	fmt.Print("Masukkan berat dan tinggi badan : ")
	fmt.Scan(&berat, &tinggi)

	BMI := berat / (tinggi * tinggi)

	fmt.Printf("BMI : %.2f\n", BMI)

}
