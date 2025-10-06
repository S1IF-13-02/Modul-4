package main

import "fmt"

func main() {
    var totalBelanja, diskon int

    fmt.Print("Masukkan total belanja: ")
    fmt.Scanln(&totalBelanja)
    fmt.Print("Masukkan diskon (%): ")
    fmt.Scanln(&diskon)

    totalAkhir := totalBelanja - (totalBelanja * diskon / 100)

    fmt.Println(totalAkhir)
}