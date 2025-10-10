package main

import "fmt"

func main() {

	var beratbadan, tinggibadan, bmi float64
	fmt.Print("masukan nilai bmi dan tinggi badan: ")
	fmt.Scan(&bmi, &tinggibadan)

	beratbadan = bmi * (tinggibadan * tinggibadan)

	fmt.Printf("%.1f", beratbadan)

}
