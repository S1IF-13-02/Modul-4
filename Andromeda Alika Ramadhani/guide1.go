package main

import "fmt"

func main() {
	var bilangan_positif int
	fmt.Scan(&bilangan_positif)

	jam := bilangan_positif / 3600
	menit := bilangan_positif / 60 % 60
	detik := bilangan_positif % 60

	fmt.Printf("%d jam %d menit %d detik", jam, menit, detik)
}
