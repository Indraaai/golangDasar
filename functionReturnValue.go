// function return value
// untuk mengembalikan nilai dari sebuah function, kita dapat menggunakan return value.
// Return value dapat digunakan untuk mengembalikan hasil dari sebuah function ke pemanggil function.
// jika sebuah function memiliki return value,
// maka kita harus menggunakan kata kunci return di dalam body function untuk mengembalikan nilai tersebut.
// untuk mengembalikan lebih dari satu nilai, kita dapat menggunakan tuple return value.

package main

import "fmt"

func getFullName(firstName, lastName string) string {
	return firstName + " " + lastName // return value berupa string yang merupakan gabungan dari firstName dan lastName
}

func getFullNameWithAge(firstName, lastName string, age int) (string, int) {
	fullName := firstName + " " + lastName
	return fullName, age // return value berupa tuple yang berisi fullName dan age
}

func main() {
	fullName := getFullName("John", "Doe")
	fmt.Println(fullName)

	fullNameWithAge, age := getFullNameWithAge("Jane", "Smith", 30)
	fmt.Printf("Full Name: %s, Age: %d\n", fullNameWithAge, age)
}
