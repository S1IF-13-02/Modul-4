package main

import (
	"fmt"
	"math"
)

func main() {
	
	var ax, ay, bx, by, cx, cy float64

	fmt.Print("Masukkan titik koordinat A: ")
	fmt.Scan(&ax, &ay)
	fmt.Print("Masukkan titik koordinat B: ")
	fmt.Scan(&bx, &by)
	fmt.Print("Masukkan titik koordinat C: ")
	fmt.Scan(&cx, &cy)

	sisiAB := math.Sqrt(math.Pow(bx - ax, 2) + math.Pow(by - ay, 2))
	sisiBC := math.Sqrt(math.Pow(cx - bx, 2) + math.Pow(cy - by, 2))
	sisiCA := math.Sqrt(math.Pow(ax - cx, 2) + math.Pow(ay - cy, 2))

	sisiTerpanjang := math.Max(sisiAB, math.Max(sisiBC, sisiCA))

	fmt.Printf("Sisi terpanjang adalah: %.2f\n", sisiTerpanjang)

}
