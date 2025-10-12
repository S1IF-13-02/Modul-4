package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func HitungSisiP(ax, ay, bx, by, cx, cy float64) float64 {
	ab := math.Sqrt(math.Pow(bx-ax, 2) + math.Pow(by-ay, 2))
	bc := math.Sqrt(math.Pow(cx-bx, 2) + math.Pow(cy-by, 2))
	ca := math.Sqrt(math.Pow(cx-ax, 2) + math.Pow(cy-ay, 2))

	return math.Max(ab, math.Max(bc, ca))
}

func main() {
	var ax, ay, bx, by, cx, cy float64

	fmt.Print("Masukan nilai Ax : ")
	fmt.Scanln(&ax)
	fmt.Print("Masukan nilai Ay : ")
	fmt.Scanln(&ay)
	fmt.Print("Masukan nilai Bx : ")
	fmt.Scanln(&bx)
	fmt.Print("Masukan nilai by : ")
	fmt.Scanln(&by)
	fmt.Print("Masukan nilai Cx : ")
	fmt.Scanln(&cx)
	fmt.Print("Masukan nilai Cy : ")
	fmt.Scanln(&cy)

	result := HitungSisiP(ax, ay, bx, by, cx, cy)
	splitResult := strings.Split(fmt.Sprintf("%.2f", result), ".")

	index1, _ := strconv.ParseFloat(splitResult[0], 64)
	index2, _ := strconv.ParseFloat(splitResult[1], 64)

	if index2 != 0 {
		fmt.Printf("%.2f\n", result)
	} else {
		fmt.Println(index1)
	}
}
