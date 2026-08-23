package main

import "fmt"

func main() {
	fmt.Println("Hello, World!", "ini adalah tipe data string")

	// fungsi yang ada dlam tipe data string

	// len untuk menghitung panjang string
	//"string"[index] untuk mengambil karakter pada indeks tertentu di dalam string
	// hasil dari "string"[index] adalah tipe data byte, sehingga perlu dikonversi ke tipe data string agar dapat ditampilkan sebagai karakter

	fmt.Println(len("Hello, World!")) // menghitung panjang string
	fmt.Println("Hello, World!"[1])   // mengambil karakter pada indeks 1
}
