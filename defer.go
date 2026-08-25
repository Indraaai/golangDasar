// fungsi defer di go adalah fungsi yang digunakan untuk menunda eksekusi sebuah statement atau fungsi hingga fungsi yang memanggilnya selesai dieksekusi.
// Fungsi defer biasanya digunakan untuk membersihkan resource, menutup file, atau melakukan tindakan lain yang harus dilakukan setelah fungsi selesai dijalankan.

package main

import "fmt"

func selesai() {
	fmt.Println("fungsi selesai dieksekusi")
}

func angka(angka1 int, angka2 int) (int, string) {
	return angka1 + angka2, "hasil penjumlahan"
}

func main() {

	defer selesai()
	result, message := angka(5, 3)
	fmt.Println(result, message) // defer akan menunda eksekusi fungsi selesai hingga fungsi main selesai dieksekusi

}
