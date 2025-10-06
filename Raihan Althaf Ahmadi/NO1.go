package main

import "fmt"

func main() {
	var total, diskon int
	fmt.Print("Masukan total belanja dan diskon (dalam persen):")
	fmt.Scan(&total)
	fmt.Scan(&diskon)
	fmt.Println(total - total*diskon/100)
}
