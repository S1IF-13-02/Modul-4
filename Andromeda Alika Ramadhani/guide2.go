package main

import "fmt"

func main() {
	var bilangan int
	fmt.Printf("Masukkan bilangan bulat positif: ")
	fmt.Scan(&bilangan)

	d1 := bilangan / 100
	d2 := bilangan / 100 % 10
	d3 := bilangan % 10

	fmt.Print(d1 <= d2 && d2 <= d3)
}
