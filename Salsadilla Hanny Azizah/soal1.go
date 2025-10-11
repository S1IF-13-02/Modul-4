package main

import (
	"fmt"
)


func main() {
	var total, diskon int
	var after int

	fmt.Print("Total Belanja = ")
	fmt.Scan(&total)
	fmt.Print("Diskon = ")
	fmt.Scan(&diskon)

	after = total - (total * diskon / 100)

	fmt.Println("Harga setelah diskon = ", after)
}