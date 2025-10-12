package main

import (
	"fmt"
)

func main() {
	var Detik int
	fmt.Print("jumlah detik: ")
	fmt.Scan(&Detik)

	jam := Detik / 3600
	menit := (Detik % 3600) / 60
	detik := Detik % 60

	fmt.Printf("%d jam %d menit %d detik\n", jam, menit, detik)

}
