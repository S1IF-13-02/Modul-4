package main

import (
	"fmt"
	"math"
)

func main() {
	var Ax, Ay, Bx, By, Cx, Cy float64

	fmt.Print("Masukkan koordinat titik Ax dan Ay: ")
	fmt.Scan(&Ax, &Ay)
	fmt.Print("Masukkan koordinat titik Bx dan By: ")
	fmt.Scan(&Bx, &By)
	fmt.Print("Masukkan koordinat titik Cx dan Cy: ")
	fmt.Scan(&Cx, &Cy)

	AB := math.Sqrt(math.Pow(Bx-Ax, 2) + math.Pow(By-Ay, 2))

	BC := math.Sqrt(math.Pow(Cx-Bx, 2) + math.Pow(Cy-By, 2))

	CA := math.Sqrt(math.Pow(Ax-Cx, 2) + math.Pow(Ay-Cy, 2))

	sisiPanjang := math.Max(AB, math.Max(BC, CA))

	fmt.Printf("Sisi Terpanjang: %.2f\n", sisiPanjang)
}
