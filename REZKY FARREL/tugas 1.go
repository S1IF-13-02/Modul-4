package main

import "fmt"

func main() {
	var totalbelanjaawal, persentasediskon float64
	fmt.Print("Masukkan total belanja awal: ")
	fmt.Scan(&totalbelanjaawal)
	fmt.Print("Masukkan persentase diskon: ")
	fmt.Scan(&persentasediskon)
	
	var totalbelanjaakhir = totalbelanjaawal * (100 - persentasediskon) / 100
	fmt.Println(totalbelanjaakhir)
}