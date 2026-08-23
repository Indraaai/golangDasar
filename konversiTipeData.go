package main

import "fmt"

func main() {
	var x int32 = 10

	// format konversi tipe data adalah dengan menggunakan tipe data yang dituju diikuti dengan nilai yang akan dikonversi
	// contoh var y adalah float64 dan x adalah int32, maka konversi tipe data dilakukan dengan menuliskan float64(x)

	var y float64 = float64(x)
	fmt.Println(y)

	// konversi tipe data dari string ke int ddan sebaliknya

	var z string = "hallo indra"
	var w = z[0] // mengambil karakter pertama dari string z

	fmt.Println(w) // hasilnya adalah tipe data byte, sehingga perlu dikonversi ke tipe data string agar dapat ditampilkan sebagai karakter

	fmt.Println(string(w)) // konversi tipe data byte ke string
}
