package main

import "fmt"

func main(){
	var totalBelanjaAwal, besarDiskon int

	fmt.Print("Masukkan total belanja: ")
	fmt.Scan(&totalBelanjaAwal)
	fmt.Print("Masukkan diskon belanja: ")
	fmt.Scan(&besarDiskon)

	diskon := totalBelanjaAwal * besarDiskon / 100
	totalBelanjaAkhir := totalBelanjaAwal - diskon

	fmt.Println("Total belanja akhir:", totalBelanjaAkhir)
	
}