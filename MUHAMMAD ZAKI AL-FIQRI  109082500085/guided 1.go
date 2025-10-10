package main

import "fmt"

func main() {
	var detik int
	var menit int
	var jam, sec int

	fmt.Print(" masukan detik : ")
	fmt.Scan(&detik)
	jam = detik / 3600
	menit = (detik / 3600) % 60
	sec = detik % 60

	fmt.Println("jam", jam)
	fmt.Println("menit ", menit)
	fmt.Println("detik ", sec)
	fmt.Printf("%d jam %d menit %d detik", jam, menit, sec)

}
