package main

import "fmt"

func main() {
	var BMI, tinggi, berat float64

	fmt.Print("Masukkan berat dan tinggi badan : ")
	fmt.Scan(&BMI, &tinggi)

	berat = BMI * tinggi * tinggi

	fmt.Printf("Berat Badan : %.f\n", berat)

}
