package main

import "fmt"

func main(){
	var hargaawal, diskon, total int
	fmt.Println("Masukkan Harga Awal: ")
	fmt.Scan(&hargaawal)
	fmt.Println("Masukkan Diskon (%): ")
	fmt.Scan(&diskon)

	total = hargaawal - (hargaawal * diskon / 100)
	fmt.Println("Total harga setelah Diskon (%):",total)

}