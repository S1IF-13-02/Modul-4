package main

import (
	"fmt"
)

func main () {
	var tb, bmi, bb float64
	fmt.Print("Masukkan BMI dan tinggi badan = " )
	fmt.Scan(&bmi, &tb)

	bb = bmi * (tb * tb)

	fmt.Printf("BMI dari hasil perhitungan adalah = %.2f\n", bb)
}