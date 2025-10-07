package main

import "fmt"

func main(){
	var dig, dig1, dig2, dig3 int
	var hasil bool

	fmt.Print("masukan angka: ")
	fmt.Scanln(&dig)

	dig1 = dig / 100
	dig2 = (dig % 100) / 10
	dig3 = dig % 10 

	hasil = dig1<=dig2 && dig2<= dig3

	fmt.Print(hasil)
	
}