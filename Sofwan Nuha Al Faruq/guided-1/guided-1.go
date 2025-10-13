package main

import "fmt"

// Fungsi utama untuk mengkonversi detik menjadi format jam, menit, dan detik
func main() {
	// Variabel untuk menyimpan input detik dari pengguna
	var d int

	// Meminta input jumlah detik dari pengguna
	fmt.Print("Detik: ")
	fmt.Scanln(&d)

	// Variabel untuk menyimpan hasil konversi jam (1 jam = 3600 detik)
	j := d / 3600

	// Variabel untuk menyimpan sisa detik setelah dikurangi jam
	sisa := d % 3600

	// Variabel untuk menyimpan hasil konversi menit (1 menit = 60 detik)
	m := sisa / 60

	// Variabel untuk menyimpan sisa detik setelah dikurangi menit
	s := sisa % 60

	// Menampilkan hasil konversi dalam format jam, menit, dan detik
	fmt.Printf("%d jam %d menit dan %d detik\n", j, m, s)
}
