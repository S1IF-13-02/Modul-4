package main

import "fmt"

func main () {
	var bilangan int
	var result bool

	fmt.Print("Masukan bilangan bulat: ")
	fmt.Scanln(&bilangan)

	d1 := bilangan / 1000
	d2 := (bilangan / 100) % 10
	d3 := (bilangan / 10) % 10

	result = d1 <= d2 && d2 <= d3
	
	fmt.Println(result)
}