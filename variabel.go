// di golang variabel hanya bisa tipe data yang sama, tidak bisa berbeda tipe data

// untuk membuat varaibel digunakna perintah var
// di golang setiap variabel harus di gunakan, jika tidak digunakan maka akan error

package main

import "fmt"

func main() {

	// inisiasi variabel di awal
	var name string

	// mengubah nilai variabel name
	name = "Dimas" // varaibel name bertipe data string

	var age int = 20           // variabel age bertipe data int
	var isMarried bool = false // variabel isMarried bertipe data bool

	fmt.Println("Nama:", name)
	fmt.Println("Umur:", age)
	fmt.Println("Status Menikah:", isMarried)

	fmt.Println("ini simbol inisiasi pertama kali :=")
	namaLengkap := "Indra Firmansyah" // inisiasi variabel dengan simbol :=
	fmt.Println("Nama Lengkap:", namaLengkap)

	// inisiasi variabel dengan simbol := dan tipe data berbeda
	umur := 25 // tipe data int
	fmt.Println("Umur:", umur)

	// deklarasi multiple variabel sekaligus
	var (
		alamat   = "Jl. Raya No. 123"
		kota     = "Jakarta"
		provinsi = "DKI Jakarta"
	)
	fmt.Println("Alamat:", alamat)
	fmt.Println("Kota:", kota)
	fmt.Println("Provinsi:", provinsi)

}
