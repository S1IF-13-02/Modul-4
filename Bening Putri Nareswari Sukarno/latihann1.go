package main

import (
	"fmt"
)

func main() {
	var time, jam, menit, detik int
	fmt.Println("waktu")
	fmt.Scan(&time)

	jam = time / 3600
	menit = time / 60 % 60
	detik = time % 60

	fmt.Println(jam, "jam", menit, "menit", detik, "detik")

}
