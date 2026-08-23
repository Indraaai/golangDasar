// > = lebih besar dari
// < = lebih kecil dari
// >= = lebih besar sama dengan
// <= = lebih kecil sama dengan
// == = sama dengan
// != = tidak sama dengan

// hasil dari perbandingan ini adalah boolean, yaitu true atau false

package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	var c bool = a > b
	fmt.Println(c) // false

	var d bool = a < b
	fmt.Println(d) // true

	var e bool = a >= b
	fmt.Println(e) // false

	var f bool = a <= b
	fmt.Println(f) // true

	var g bool = a == b
	fmt.Println(g) // false

	var h bool = a != b
	fmt.Println(h) // true

}
