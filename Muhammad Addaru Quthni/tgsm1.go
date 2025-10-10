package main
import (
	"fmt"
)
func main(){
	var nilaiBMI, tinggiBadan float64

	fmt.Print("Masukkan nilai BMI : ")
	fmt.Scan(&nilaiBMI)
	fmt.Print("Masukkan tinggi badan : ")
	fmt.Scan(&tinggiBadan)

	beratBadan := nilaiBMI * (tinggiBadan * tinggiBadan)
	fmt.Printf("Berat badan adalah: %.f kg\n", beratBadan)
}