package main

import "fmt"
func main(){
	var beratbadan float64
	fmt.Println("masukan berat badan =")
	fmt.Scanln(&beratbadan)

	var tinggibadan float64
	fmt.Println("masukan tinngi badan =")
	fmt.Scanln(&tinggibadan)

	bmi := beratbadan * (tinggibadan * tinggibadan)

	fmt.Printf("bmi anda = %.0f\n",bmi)
}