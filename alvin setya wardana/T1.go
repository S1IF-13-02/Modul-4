package main

import "fmt"

func main() {
	var awal, setelahdiskon int
	var diskon, totaldinkon int

	fmt.Print("Masukkanharga barang: ")
	fmt.Scan(&awal)
	fmt.Print("Diskon: ")
	fmt.Scan(&diskon)

	totaldinkon = (awal * diskon) / 100
	setelahdiskon = awal - totaldinkon

	fmt.Println("Harga setelah diskon: ", setelahdiskon)
}
