// percabngan dengan if else statement
package main

import "fmt"

func main() {
	x := 10
	// percabangan di lakukan dengan operator boolean, dimana jika kondisi bernilai true maka akan mengeksekusi blok kode if, jika kondisi bernilai false maka akan mengeksekusi blok kode else
	if x > 5 {
		fmt.Println("x lebih besar dari 5")
	} else {
		fmt.Println("x tidak lebih besar dari 5")
	}

	// else if statement digunakan untuk menambahkan kondisi tambahan pada percabangan if else, dimana jika kondisi pertama bernilai false maka akan mengecek kondisi kedua, jika kondisi kedua bernilai true maka akan mengeksekusi blok kode else if, jika kondisi kedua bernilai false maka akan mengeksekusi blok kode else
	nilai := 75
	if nilai >= 90 {
		fmt.Println("Nilai A")
	} else if nilai >= 80 {
		fmt.Println("Nilai B")
	} else {
		fmt.Println("Nilai C")
	}

	// if short statement digunakan untuk mengeksekusi blok kode if dengan kondisi yang lebih sederhana, dimana jika kondisi bernilai true maka akan mengeksekusi blok kode if, jika kondisi bernilai false maka akan mengeksekusi blok kode else
	if hasil := nilai * 2; hasil > 150 {
		fmt.Println("Hasil lebih besar dari 150")
	} else {
		fmt.Println("Hasil tidak lebih besar dari 150")
	}
}
