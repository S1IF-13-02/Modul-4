package main

import "fmt"

func main() {

	var bmi, tinggi, berat float64

	fmt.Print("masukan nilai bmi : ")
	fmt.Scan(&bmi)
	fmt.Print("masukan tinggi badan : ")
	fmt.Scan(&tinggi)

	berat = bmi * (tinggi * tinggi)

	fmt.Printf("%.0f", berat)
}
