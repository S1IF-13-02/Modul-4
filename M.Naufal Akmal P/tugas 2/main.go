package main

import (
	"fmt"
	"math"
)

func main() {
	
	var bmi, tinggi float64
	
	fmt.Print("masukan bmi: ")
	fmt.Scan(&bmi)
	fmt.Print("masukan tinggi: ")
	fmt.Scan(&tinggi)

	beratBadan := bmi * (tinggi * tinggi)
	hasilBulat := math.Round(beratBadan)

	fmt.Println(int(hasilBulat))
}