package main

import "fmt"

func main() {

	var hargaawal, hargasetelahdiskon, diskon, totaldiskon int

	fmt.Print("masukan harga awal : ")
	fmt.Scan(&hargaawal)
	fmt.Print("masukan diskon :")
	fmt.Scan(&diskon)

	totaldiskon = (hargaawal * diskon) / 100
	hargasetelahdiskon = hargaawal - totaldiskon

	fmt.Print(hargasetelahdiskon)

}
