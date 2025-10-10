package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("masukan angka : ")
	fmt.Scan(&bilangan)

	angka1 := bilangan / 100
	angka2 := (bilangan % 100) / 10
	angka3 := bilangan % 10

	fmt.Print(angka1 <= angka2 && angka2 <= angka3)

}
