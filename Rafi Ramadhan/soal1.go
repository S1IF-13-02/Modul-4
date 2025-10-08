package main

import "fmt"

func main() {
    var totalBelanja int
    var diskonPersen int

    fmt.Print("Masukkan total belanja: ")
    fmt.Scan(&totalBelanja)

    fmt.Print("Masukkan besar diskon (%): ")
    fmt.Scan(&diskonPersen)
    diskon := float64(totalBelanja) * float64(diskonPersen) / 100

    totalAkhir := float64(totalBelanja) - diskon

    fmt.Printf("Total belanja setelah diskon: %.0f\n", totalAkhir)
}
