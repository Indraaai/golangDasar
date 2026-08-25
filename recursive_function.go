// recursive function in go

// recursive funtion di go adalah function yang memanggil dirinya sendiri,
// recursive function biasanya digunakan untuk menyelesaikan masalah yang bisa dipecah menjadi sub masalah yang lebih kecil

// contoh recursive function adalah menghitung faktorial dari sebuah angka

package main

import "fmt"

func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * faktorial(n-1)
}

func main() {
	fmt.Println(faktorial(5)) // output: 120
}
