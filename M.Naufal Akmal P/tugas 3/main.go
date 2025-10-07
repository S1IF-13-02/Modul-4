package main

import (
	"fmt"
	"math"
)
type Point struct {
	X, Y float64
}
func distance(p1, p2 Point) float64 {
	
	deltaX := p2.X - p1.X
	deltaY := p2.Y - p1.Y

	return math.Sqrt(math.Pow(deltaX, 2) + math.Pow(deltaY, 2))
}

func main() {
	var a, b, c Point
	fmt.Println("masukan a.X dan a.Y: ")
	fmt.Scanln(&a.X, &a.Y)
	fmt.Println("masukan b.X dan b.Y: ")
	fmt.Scanln(&b.X, &b.Y)
	fmt.Println("masukan c.X dan c.Y: ")
	fmt.Scanln(&c.X, &c.Y)

	sisiAB := distance(a, b)
	sisiBC := distance(b, c)
	sisiAC := distance(a, c)

	sisiTerpanjang := math.Max(sisiAB, math.Max(sisiBC, sisiAC))

	fmt.Printf("%.2f\n", sisiTerpanjang)
}