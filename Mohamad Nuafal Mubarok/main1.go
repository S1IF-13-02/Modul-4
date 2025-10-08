package main

import "fmt"

func main() {
    var totalBelanja int
    var diskon int

    fmt.Print("totalBelanja:")
    fmt.Scanln(&totalBelanja)

    fmt.Print("diskon:")
    fmt.Scanln(&diskon)

    totalAkhir := totalBelanja - (totalBelanja * diskon / 100)

    fmt.Println(totalAkhir)
}
