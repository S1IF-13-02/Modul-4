package main
import "fmt"

func main() {
	var BMI, berat, tinggi float32
	fmt.Print("Masukan berat badan : ")
	fmt.Scan(&berat)

	fmt.Print("Masukan tinggi badan : ")
	fmt.Scan(&tinggi)

	BMI = berat / (tinggi * tinggi)

	fmt.Printf("Nilai BMI orang tersebut adalah : %.2f\n", BMI)
}