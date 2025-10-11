package main

import "fmt"

func main() {
	var BMI float64
	var TB float64
	var BB float64
	fmt.Printf("Masukkan BMI: ")
	fmt.Scan(&BMI)
	fmt.Printf("Masukkan TB: ")
	fmt.Scan(&TB)
	BB = BMI * (TB * TB)
	fmt.Printf("Nilai BB: %d", int(BB))

}
