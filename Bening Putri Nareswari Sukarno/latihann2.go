package main

import (
	"fmt"
)

func main() {
	var nilai, nilai1, nilai2, nilai3 int
	var hasil bool
	fmt.Println("Masukkan angka (100-999): ")
	fmt.Scanln(&nilai)

	nilai1 = nilai / 100
	nilai2 = (nilai % 100) / 10
	nilai3 = nilai % 10

	hasil = nilai1 <= nilai2 && nilai2 <= nilai3

	fmt.Println("Apakah digitnya berurutan?", hasil)

}
