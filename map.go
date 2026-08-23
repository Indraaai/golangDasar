// map adlah tipe data dalam golang dengan key dan value, dimana key bersifat unik dan value bisa berupa tipe data apapun.
// Map digunakan untuk menyimpan pasangan key-value dan memungkinkan akses cepat ke nilai berdasarkan kunci.
// format penulisan map: map[keyType]valueType

package main

import "fmt"

func main() {
	// contoh deklarasi map dengan key bertipe string dan value bertipe int
	Umur := map[int]string{
		25: "Alice",
		30: "Bob",
		35: "Charlie",
	}
	fmt.Println(Umur)
	// akses nilai dalam map menggunakan key
	fmt.Println("Nama pada umur 30 adalah:", Umur[30])

	// contoh deklarasi map dengan key bertipe string dan value bertipe int
	Nilai := map[string]int{
		"Matematika": 90,
		"Fisika":     85,
		"Kimia":      95,
	}
	fmt.Println(Nilai)
	fmt.Println("Nilai Matematika adalah:", Nilai["Matematika"])
	fmt.Println("Nilai Fisika adalah:", Nilai["Fisika"])
	fmt.Println("Nilai Kimia adalah:", Nilai["Kimia"])

	// fungsi yang ada dalam map
	//len() digunakan untuk menghitung jumlah elemen dalam map
	fmt.Println("Jumlah elemen dalam map Nilai adalah:", len(Nilai)) // hasilnya 3

	// mengambil data dalam map menggunakan key
	fmt.Println("Nilai Matematika adalah:", Nilai["Matematika"]) // hasilnya 90

	// menambahkan data baru ke dalam map
	Nilai["Biologi"] = 80
	fmt.Println("Nilai Biologi adalah:", Nilai["Biologi"]) // hasilnya 80

	// mengubah nilai dalam map menggunakan key
	Nilai["Fisika"] = 95
	fmt.Println("Nilai Fisika setelah diubah adalah:", Nilai["Fisika"]) // hasilnya 95

	// menghapus data dalam map menggunakan key
	delete(Nilai, "Kimia")
	fmt.Println("Nilai Kimia setelah dihapus adalah:", Nilai["Kimia"]) // hasilnya 0 karena data sudah dihapus

	// membuat map baru menggunakan make()
	// contoh deklarasi map dengan key bertipe string dan value bertipe int menggunakan make()
	Nilai2 := make(map[string]int)
	Nilai2["Bahasa Indonesia"] = 85
	Nilai2["Bahasa Inggris"] = 90
	fmt.Println(Nilai2)
	fmt.Println("Nilai Bahasa Indonesia adalah:", Nilai2["Bahasa Indonesia"]) // hasilnya 85
	fmt.Println("Nilai Bahasa Inggris adalah:", Nilai2["Bahasa Inggris"])     // hasilnya 90
}
