// perulangan di go menggunakan statement for, dimana statement for memiliki 3 bagian yaitu inisialisasi, kondisi, dan post statement.
// Inisialisasi digunakan untuk mendeklarasikan variabel yang akan digunakan pada perulangan, kondisi digunakan untuk menentukan kapan perulangan akan berhenti, dan post statement digunakan untuk mengubah nilai variabel yang digunakan pada perulangan.
package main

import "fmt"

func main() {

	angka := 1 // inisiasi variabel awal

	for angka <= 5 { // kondisi perulangan
		angka++ // post statement
		fmt.Println("perulangan angka:", angka)
	}
	fmt.Println("selesai")

	// inisiasi for dengan lebih sederhana
	for counter := 1; counter <= 5; counter++ { // inisialisasi, kondisi, dan post statement
		fmt.Println("perulangan counter:", counter)
	}

	// perulangan dengan menggunakan range, dimana range digunakan untuk mengiterasi elemen dari sebuah array, slice, map, atau string.
	// range akan mengembalikan 2 nilai yaitu index dan value dari elemen yang sedang diiterasi.
	// jika hanya ingin menggunakan value saja, maka bisa menggunakan _ (underscore) untuk mengabaikan index.
	// jika hanya ingin menggunakan index saja, maka bisa menggunakan _ (underscore) untuk mengabaikan value.

	// contoh perulangan dengan menggunakan range pada array
	angkaArray := [5]int{1, 2, 3, 4, 5}    // inisialisasi array
	for index, value := range angkaArray { // inisialisasi index dan value
		fmt.Println("perulangan index:", index, "value:", value)
	}

	// contoh perulangan dengan menggunakan range pada slice
	angkaSlice := []int{1, 2, 3, 4, 5}
	for index, value := range angkaSlice {
		fmt.Println("perulangan index:", index, "value:", value)
	}

	// contoh perulangan dengan menggunakan range pada map
	angkaMap := map[string]int{"satu": 1, "dua": 2, "tiga": 3}
	for key, value := range angkaMap {
		fmt.Println("perulangan key:", key, "value:", value)
	}

	// contoh perulangan dengan menggunakan range pada string
	angkaString := "12345"
	for index, value := range angkaString {
		fmt.Println("perulangan index:", index, "value:", string(value))
	}

	// contoh perulangan dengan tanda _ (underscore) untuk mengabaikan index
	for _, value := range angkaArray {
		fmt.Println("perulangan value:", value)
	}
}
