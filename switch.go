// switch expression

package main

import "fmt"

func main() {
	// switch statement digunakan untuk mengecek nilai dari sebuah variabel atau ekspresi, dimana jika nilai dari variabel atau ekspresi sesuai dengan salah satu case maka akan mengeksekusi blok kode case tersebut, jika tidak ada case yang sesuai maka akan mengeksekusi blok kode default
	nilai := 75
	switch nilai {
	case 90: // kondisi 1
		fmt.Println("Nilai A")
	case 80: // kondisi 2
		fmt.Println("Nilai B")
	default: // default case digunakan untuk mengeksekusi blok kode jika tidak ada case yang sesuai dengan nilai dari variabel atau ekspresi
		fmt.Println("Nilai C")
	}

	// switch short statement digunakan untuk mengecek nilai dari sebuah variabel atau ekspresi dengan kondisi yang lebih sederhana, dimana jika nilai dari variabel atau ekspresi sesuai dengan salah satu case maka akan mengeksekusi blok kode case tersebut, jika tidak ada case yang sesuai maka akan mengeksekusi blok kode default
	switch hasil := nilai * 2; hasil {
	case 150: // kondisi 1
		fmt.Println("Hasil A")
	case 100: // kondisi 2
		fmt.Println("Hasil B")
	default: // default case digunakan untuk mengeksekusi blok kode jika tidak ada case yang sesuai dengan nilai dari variabel atau ekspresi
		fmt.Println("Hasil C")
	}
}
