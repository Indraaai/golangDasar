// break dan continue digunakan untuk mengontrol alur perulangan,
// dimana break digunakan untuk menghentikan perulangan dan keluar dari blok perulangan,
// sedangkan continue digunakan untuk melewati iterasi saat ini dan melanjutkan ke iterasi berikutnya.

package main

import "fmt"

func main() {

	// break statement

	for i := 1; i <= 10; i++ {
		if i == 5 {
			break // keluar dari perulangan ketika i sama dengan 5
		}
		fmt.Println("perulangan i:", i)
	}

	// continue statement

	for j := 1; j <= 10; j++ {
		if j%2 == 0 {
			continue // melewati iterasi ketika j adalah bilangan genap
		}
		fmt.Println("perulangan j:", j)
	}
}
