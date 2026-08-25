// panic fungsi adalah fungsi yang digunakan untuk menghentikan eksekusi program secara paksa ketika terjadi kondisi yang tidak diinginkan atau error yang tidak dapat ditangani.
// Fungsi panic biasanya digunakan dalam situasi kritis di mana program tidak dapat melanjutkan eksekusi dengan aman.

package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(err bool) {
	defer endApp() // defer akan menunda eksekusi fungsi endApp hingga fungsi runApp selesai dieksekusi
	if err {
		panic("Terjadi error kritis") // panic akan menghentikan eksekusi program dan men
	}
}

func main() {
	runApp(true)
}
