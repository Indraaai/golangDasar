// Struct adalah blueprint untuk membuat object dalam go
/* cara menulis struct adalah
   type NamaStruct Struct {
	   field1 type1
	   field2 type2
   }*/

package main

import "fmt"

type User struct {
	Name    string
	Address string
	Age     uint8
}

// Struct Literal
// cara membuat struct literal adalah
/* namaVariabel :=  User{
   Name: "Budi",
   Address: "Jakarta",
   Age: 25,
}*/

func main() {

	// deklarasi Struct ke object dengan nama user
	var user User
	// mengisi nilai field dari struct user dengan nama
	user.Name = "Budi"
	user.Address = "Jakarta"
	user.Age = 25

	fmt.Println(user)
	fmt.Println(user.Name)
	fmt.Println(user.Address)
	fmt.Println(user.Age)

}
