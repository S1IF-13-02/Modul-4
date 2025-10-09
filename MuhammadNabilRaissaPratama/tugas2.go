package main

import "fmt"
func main(){
	var bilanganbulat int
	fmt.Println("masukan angka:")
	fmt.Scanln(&bilanganbulat)

	d1 := bilanganbulat / 100
	d2 := bilanganbulat % 10 
	d3 := bilanganbulat % 100 
	fmt.Print(d1<=d2&& d2 <=d3)
   

}