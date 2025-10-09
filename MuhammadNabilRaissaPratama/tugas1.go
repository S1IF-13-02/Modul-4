package main

import "fmt"
func main() {
	var totaldetik int
	fmt.Println("masukan detik:")
	fmt.Scanln(&totaldetik)

	jam := totaldetik / 3600
	sisa := totaldetik % 3600
    menit := sisa / 60
	detik := sisa % 60
	fmt.Printf("%d jam %d menit %d detik", jam, menit, detik)
}
