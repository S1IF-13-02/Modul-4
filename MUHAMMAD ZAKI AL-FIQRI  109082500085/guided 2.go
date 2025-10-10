package main

import "fmt"

func main() {

	var angka1, angka2, angka3, bilangan int

	fmt.Print("masukan bilangan : ")
	fmt.Scan(&bilangan)

	angka1 = bilangan / 100
	angka2 = (bilangan % 100) / 10
	angka3 = angka2 % 10

	fmt.Print(angka1 <= angka2 && angka2 <= angka3)

}
