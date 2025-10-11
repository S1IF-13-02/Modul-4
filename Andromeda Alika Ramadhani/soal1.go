package main

import "fmt"

func main() {
	var total, diskon int
	fmt.Print("Total Belanjaan: ")
	fmt.Scan(&total)
	fmt.Print("Diskon (dalam persen): ")
	fmt.Scan(&diskon)

	Total_Belanja := total - (total * diskon / 100)

	fmt.Printf("Total belanja anda adalah %d", Total_Belanja)
}
