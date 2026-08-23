// type decalaration digunakan untuk mendeklarasikan tipe data baru di Go.
// Tipe data ini bisa berupa struct, interface, atau alias tipe data.
// Berikut adalah beberapa contoh penggunaan type declaration di Go:

package main

import "fmt"

func main() {

	// mendeklarasikan tipe data baru bernama "nama" yang merupakan alias dari tipe data string
	type nama string

	// contoh penggunaan tipe data baru "nama"
	var myName nama = "Dimas"
	fmt.Println(myName)

	// mendeklarasikan tipe data baru bernama "umur" yang merupakan alias dari tipe data int
	type umur int

	// contoh penggunaan tipe data baru "umur"
	var myAge umur = 25
	fmt.Println(myAge)
}
