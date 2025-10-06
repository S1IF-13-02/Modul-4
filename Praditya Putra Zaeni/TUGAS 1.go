package main

import "fmt"

func main() {
	var totalAwal, diskonPersen int

	fmt.Scan(&totalAwal)
	fmt.Scan(&diskonPersen)

	potongan := totalAwal * diskonPersen / 100
	totalAkhir := totalAwal - potongan
	fmt.Println(totalAkhir)
}
