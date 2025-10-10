package main

import "fmt"

func main() {

	var semuadetik int

	fmt.Print("masukan detik : ")
	fmt.Scan(&semuadetik)

	jam := semuadetik / 3600
	sisa := semuadetik % 3600
	menit := sisa / 60
	detik := sisa % 60

	fmt.Printf("%d jam %d menit dan %d detik", jam, menit, detik)
}
