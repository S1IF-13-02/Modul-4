package main

import "fmt"

func main() {
	var awal, diskon, akhir float64

	fmt.Print("Total belanja awal: ")
	fmt.Scanln(&awal)
	fmt.Print("Besar diskon (%): ")
	fmt.Scanln(&diskon)
	fmt.Printf("(diskon = %.0f%%)\n", diskon)
	akhir = awal - (awal * diskon / 100)
	fmt.Println("Total belanja akhir =", akhir)
}
