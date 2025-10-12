package main

import "fmt"

func main() {
	var totalBelanjaAwal, diskonPersen int

	fmt.Scan(&totalBelanjaAwal)
	fmt.Scan(&diskonPersen)

	diskonNominal := (totalBelanjaAwal) * (diskonPersen) / 100.0
	totalBelanjaAkhir := (totalBelanjaAwal) - diskonNominal

	fmt.Println(totalBelanjaAkhir)
}
