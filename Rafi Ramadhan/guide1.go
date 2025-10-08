package main
 
import "fmt"

func main() {
	var d int

	fmt.Print("Detik: ")
	fmt.Scanln(&d)
	 j := d / 3600
	 sisa := d % 3600
	 m := sisa / 60
	 s := sisa % 60
	 fmt.Printf("%d jam %d menit dan %d detik\n", j, m, s)
}
