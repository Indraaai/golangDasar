// tipe data slice adalah tipe data yang bisa menampung banyak data dengan tipe data yang sama
// tipe data slice bisa diubah isinya, bisa ditambah, dikurangi, dan dihapus
// tipe data slice memiliki 3 data, yaitu length, capacity, dan pointer
// index slice dimulai dari 0, dan index terakhir adalah length - 1
// lenght adalah jumlah data yang ada di dalam slice
// pointer adalah alamat dari data pertama di dalam slice

package main

import "fmt"

func main() {

	// array declare
	array := [...]string{"indra", "budi", "caca", "didi", "erik"}

	// slice declare
	slice := array[1:4] // slice dari array, dimulai dari index 1 sampai index 4 (tidak termasuk index 4)

	fmt.Println(slice)

	slice2 := array[2:] // slice dari array, dimulai dari index 2 sampai akhir array

	fmt.Println(slice2)

	slice3 := array[:3] // slice dari array, dimulai dari index 0 sampai index 3 (tidak termasuk index 3)

	fmt.Println(slice3)

	slice4 := array[:] // slice dari array, dimulai dari index 0 sampai akhir array

	fmt.Println(slice4)

	// declare slice langsung tanpa array
	slice5 := []string{"indra", "budi", "caca", "didi", "erik"}

	fmt.Println(slice5)

	// fungsi fungsi slice
	fmt.Println("length slice5:", len(slice5))   // panjang slice5
	fmt.Println("capacity slice5:", cap(slice5)) // kapasitas slice5

	// fungsi append untuk menambah data ke dalam slice
	slice5 = append(slice5, "fajar") // menambah data "fajar" ke dalam slice5
	fmt.Println(slice5)

	slice6 := slice5[1:4] // slice dari slice5, dimulai dari index 1 sampai index 4 (tidak termasuk index 4)
	fmt.Println(slice6)

	slice6[0] = "budi2" // mengubah data di index 0 dari slice6
	fmt.Println(slice6)
	fmt.Println(slice5) // data di slice5 juga berubah karena slice6 adalah referensi dari slice5

	// membuat slice dengan fungsi make
	slice7 := make([]string, 3, 5) // membuat slice dengan panjang 3 dan kapasitas 5
	fmt.Println(slice7)
	// menambahkan data ke dalam slice7
	slice7[0] = "indra"
	slice7[1] = "budi"
	slice7[2] = "caca"
	fmt.Println(slice7)
	fmt.Println("length slice7:", len(slice7))
	fmt.Println("capacity slice7:", cap(slice7))

	// untuk menambahakn data ke slice 7 di di luar length slice7, maka harus menggunakan fungsi append
	slice7 = append(slice7, "didi")
	slice7 = append(slice7, "erik")

	// jika panjang slice lebih dari kapasitasnya, maka kapasitas slice akan bertambah dua kali lipat

	slice7 = append(slice7, "fajar")
	fmt.Println(slice7)
	fmt.Println("length slice7:", len(slice7))
	fmt.Println("capacity slice7:", cap(slice7))

	// fungsi copy untuk menyalin data dari satu slice ke slice lain
	slice8 := make([]string, 6)
	copy(slice8, slice7)
	fmt.Println("ini slice8:", slice8)
}
