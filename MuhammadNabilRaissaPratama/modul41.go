package main

import "fmt"

func main(){

	var totalbelanja int
	var totaldiskon int

	fmt.Printf("total belanja :")
	fmt.Scanln(&totalbelanja)

	fmt.Printf("total diskon :")
    fmt.Scanln(&totaldiskon)

	totalakhir := totalbelanja - (totalbelanja * totaldiskon / 100)

	fmt.Println(totalakhir)
}