package main

import "fmt"

func main(){
	var detikk int
	fmt.Println("Berapa detik:")
	fmt.Scan(&detikk)
	jam := (detikk / 3600)
	menit := (detikk % 3600 / 60)
	detik := (detikk % 60)
	fmt.Println(jam, "Jam", menit, "Menit dan", detik, "Detik")
} 