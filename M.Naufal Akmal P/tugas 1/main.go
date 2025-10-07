package main

import "fmt"

func main() {
	
	var totalBelanjaAwal, persentaseDiskon int
	
	fmt.Print("masukan total belanja awal: ")
	fmt.Scanln(&totalBelanjaAwal)
	fmt.Print("masukan persentase diskon: ")
	fmt.Scanln(&persentaseDiskon)

	
	var totalBelanjaAkhir = totalBelanjaAwal * (100 - persentaseDiskon) / 100
	fmt.Println(totalBelanjaAkhir)
}