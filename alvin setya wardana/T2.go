package main

import "fmt"

func main() {
	var bmi, berat, tinggi float64

	fmt.Print("masukan bmi: ")
	fmt.Scanln(&bmi)
	fmt.Print("masukan tinggi: ")
	fmt.Scanln(&tinggi)
	berat = bmi * tinggi * tinggi
	fmt.Printf("berat badan: %.0fkgn\n", berat)
}

