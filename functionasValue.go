// function sebagai nilai
// function as value adalah konsep dimana function dapat disimpan dalam variabel,
// dikirim sebagai argumen ke function lain, atau dikembalikan dari function lain.
// Ini memungkinkan kita untuk menggunakan function dengan cara yang lebih fleksibel dan dinamis.

package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	// Menyimpan function dalam variabel
	var operation func(int, int) int = add
	result := operation(5, 3)
	fmt.Println("Result:", result)
}
