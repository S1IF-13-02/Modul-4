package main

import "fmt"

func main(){
	var bmi, tinggibadan, beratbadan float64 
	fmt.Scan(&bmi, &tinggibadan)

	beratbadan = bmi * (tinggibadan * tinggibadan)
	fmt.Printf("Berat Badan: %.0f", beratbadan)

}