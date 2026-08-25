// method di go adalah sebuah fungsi yang dimiliki oleh Struct
// format penulisan method adalah
/*
func (namaVariabel NamaStruct) NamaMethod(){
	// kode
}*/

// method juga bisa menerima paramaeter dan return value seperti function biasa

package main

import "fmt"

type Pelanggan struct {
	Name string
	Age  int
}

// contoh sebuah metod yang mengimplementasikan Struct  Pelanggan

func (u Pelanggan) user() {
	fmt.Println("hello user dengan nama", u.Name)
}

func (u Pelanggan) Umur() {
	fmt.Println("hallo", u.Name, "umur kamu adalah ", u.Age)
}

func main() {
	sapa := Pelanggan{Name: "indra", Age: 21}
	sapa.user()
	sapa.Umur()

}
