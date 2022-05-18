package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

// short notation for contactInfo
// type person struct {
// 	firstName string
// 	lastName  string
// 	contactInfo
// }

func main() {
	// simple struct: firstName, lastName
	// person1 := person{"John", "Doe"}
	// person2 := person{firstName: "John", lastName: "Doe"}

	// fmt.Println(person1)
	// fmt.Println(person2)

	// // default values for a new instance
	// var person3 person
	// fmt.Println(person3)

	// // key-value format
	// fmt.Printf("%+v\n", person3)

	// person3.firstName = "Johnny"
	// fmt.Printf("%+v\n", person3)

	// fmt.Println("\n")

	// nested struct, with contactInfo
	person1 := person{
		firstName: "John",
		lastName:  "Doe",
		contact: contactInfo{
			email:   "john@doe.com",
			zipCode: 90000,
		},
	}

	person1.updateName("Johnny")
	person1.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

// print function can be called on any variable of type person
func (p person) print() {
	fmt.Printf("%+v\n", p)
}
