// parameter function
// nama parameter function bersifat unik dan tidak boleh sama dengan nama parameter function lainnya dalam satu blok function yang sama.
// Parameter function digunakan untuk menerima input dari pemanggil function dan dapat digunakan di dalam body function.

package main

import "fmt"

func sayHelloIndra(hello, name string) {
	fmt.Println("hello", hello, "my name is", name)
}

func main() {

	sayHelloIndra("indra", "Dimas")
}
