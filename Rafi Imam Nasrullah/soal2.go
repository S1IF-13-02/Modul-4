package main

import 

"fmt"


func main() {
	var BMI, tinggi, berat float64

	fmt.Print("Masukkan nilai BMI: ")
	fmt.Scanln(&BMI)

	fmt.Print("Masukkan tinggi badan (meter): ")
	fmt.Scanln(&tinggi)

	berat = BMI * (tinggi * tinggi)

	fmt.Printf("Berat badan adalah: %.0f\n", berat)
	
}
