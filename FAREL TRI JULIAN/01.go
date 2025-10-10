package main

import "fmt"

func main() {

	var awal, setelahdiskon, diskon, totaldiskon int

	fmt.Print("masukan harga : ")
	fmt.Scan(&awal)
	fmt.Print("Diskon : ")
	fmt.Scan(&diskon)

	totaldiskon = (awal * diskon) / 100
	setelahdiskon = awal - totaldiskon

	fmt.Print(setelahdiskon)
}
