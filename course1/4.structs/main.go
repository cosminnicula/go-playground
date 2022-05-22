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

	// pass by value
	person1.updateName("Johnny")
	person1.print() // prints John, because receiver parameters are passed as values

	// pass by reference via pointer
	person1Pointer := &person1
	person1Pointer.updateNameByReference("Johnny")
	person1.print()

	// pointer shortcut

	// although  updateNameByReference expects the receiver to be of type pointer
	// to person, Go automatically coverts the person1 variable of type person into
	// a pointer to person
	person1.updateNameByReference("John")
	person1.print()

	//however, there are both reference and value types:
	//    - structs are value types
	//    - slices are structs that contain 3 fields: lenght, capacity, and a pointer to the head of the array.
	//		  when slices are passed as arguments to a function, the entire (length, capacity, pointer to head) structure
	//      is copied over, which means that will point to the same array in memory.
	//value types: int, float, string, bool, struct
	//reference types: slices, maps, channels, pointers, functions
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (pointerToPerson *person) updateNameByReference(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// print function can be called on any variable of type person
func (p person) print() {
	fmt.Printf("%+v\n", p)
}
