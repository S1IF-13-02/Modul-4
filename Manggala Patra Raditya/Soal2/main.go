package main

import (
	"fmt"
	"math"
)

func main() {
	var bmi, t float64
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bmi, &t)
	fmt.Printf("hasil: %.0f kg\n", math.Round(bmi*t*t))
}
