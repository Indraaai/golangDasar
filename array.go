// tipe data array adalah tipe data yang bisa menampung banyak data dengan tipe data yang sama
// array memiliki index yang dimulai dari 0
// array memiliki panjang yang tetap, tidak bisa diubah dan harus ditentukan saat deklarasi array
// index array bisa diakses dengan menggunakan tanda kurung siku []
// index array dimulai dari 0, jadi jika panjang array adalah 5, maka index array adalah 0, 1, 2, 3, 4
// cara mendeklarasikan array adalah dengan menggunakan tanda kurung siku [] dan menentukan panjang array, kemudian diikuti dengan tipe data array
package main

import "fmt"

func main() {

	var array [5]int // inisiasi array dengan panjang 5 dan tipe data int
	// mengisi array dengan data
	array[0] = 10
	array[1] = 20
	array[2] = 30
	array[3] = 40
	array[4] = 50

	fmt.Println(array)    // [10 20 30 40 50]
	fmt.Println(array[0]) // 10

	//membuat array dengan langsung mengisi data
	var array2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2) // [1 2 3 4 5]

	// fungsi yang ada di array

	// fungsi len() untuk mengetahui panjang array
	fmt.Println(len(array)) // 5
	// fungsi index untuk mengetahui index array
	fmt.Println(array[0])
	// fungsi array[index] = value untuk mengubah nilai array
	array[0] = 100
	fmt.Println(array[0]) // 100
	// fungsi cap() untuk mengetahui kapasitas array
	fmt.Println(cap(array)) // 5
	// fungsi copy() untuk menyalin array
	var array3 [5]int
	copy(array3[:], array[:])
	fmt.Println(array3) // [10 20 30 40 50]

	// jika tidak ingin menentukan panjang array, bisa menggunakan ... untuk menentukan panjang array secara otomatis
	var array4 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(array4) // [1 2 3 4 5]
}
