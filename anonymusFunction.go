// anonymus function adalah function tanpa nama yang bisa dibuat langsung dalam variabel atau parameter

package main

import "fmt"

// deklarasi type fungsi sebagai function yang berparameter string dan return bool
type Fungsi func(string) bool

type cekbagi func(int) bool

// fungsi blaklist memiliki 2 paramter nama = string dan blk dalah type Type fungsi
func Blaklist(name string, blk Fungsi) {
	if blk(name) { // blk mnegecek apakah paramter name true
		fmt.Println("kamu di block")
	} else {
		fmt.Println("kamu aman")
	}
}

func hasilBagi(angka int, angkaHabis cekbagi) {
	if angkaHabis(angka) {
		fmt.Println("angka anda genap")
	} else {
		fmt.Println("angka anda ganjil")
	}
}

// membuat fungsi

func main() {
	banned := func(name string) bool { // banned adalah varaiabel anonymus function yang memiliki kontrak yangsama dengan Type fungsi di atas yang secara langsung sama seperti paramter blk di funtion blacklist
		return name == "blok" // fungsi ini mereturn parameter name == string "blok" yan artinya jika paramter yang di kirim ke blakslit namenya sama dengan paramter fungsi maka pengecke fungsi blaklist bernilai true
	}

	Blaklist("indra", banned) //

	tesangka := func(angka int) bool {
		return angka%2 == 0
	}

	hasilBagi(10, tesangka)
}
