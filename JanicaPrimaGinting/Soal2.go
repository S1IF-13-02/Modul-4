package main 

import "fmt"

func main(){
	var bilanganpositif int
	var d1,d2,d3 int
	fmt.Println("Masukkan bilangan positif:")
	fmt.Scan(&bilanganpositif)

	d1 = (bilanganpositif / 100)
	d2 = (bilanganpositif % 100 / 10)
	d3 = (bilanganpositif % 10)
	fmt.Print(d1<=d2 && d2<=d3)
}