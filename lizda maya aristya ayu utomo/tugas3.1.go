package main

import (
	"fmt"
)

func main() {
	var totalawal int
	var diskonpersen int

	fmt.Print("masukan total awal: ")
	fmt.Scanln(&totalawal)

	fmt.Print("masukan diskon: ")
	fmt.Scanln(&diskonpersen)

	jumlahdiskon := float64(diskonpersen) / 100 * float64(totalawal)

	totalakhir := float64(totalawal) - jumlahdiskon

	fmt.Printf("total belanja setelah diskon: %d\n", int(totalakhir))

}
