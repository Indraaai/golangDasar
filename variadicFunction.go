// variadic function.go
// variadic function disimbolkan dengan tanda ... (titik tiga) sebelum tipe data parameter terakhir.
// variadic function dapat menerima nol atau lebih argumen dari tipe data yang sama.

package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total

}

func main() {
	result := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum:", result)
}
