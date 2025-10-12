package main

import "fmt"

func main() {
	var konversi, jam, menit, detik int

	fmt.Print("Masukkan detik : ")
	fmt.Scanln(&konversi)

	jam = konversi / 3600
	menit = (konversi % 3600) / 60
	detik = menit % 60

	fmt.Println(jam, "jam", menit, "menit", detik, "detik")
}
