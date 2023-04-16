package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	var alex person
	jim := person{firstName: "Jim", lastName: "Party", contactInfo: contactInfo{email: "hello@gmail.com", zipCode: 94000}}
	// if left alex is person, it will return as empty string
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.print()
	jim.print()

	// Pointer
	// &variable = give me memory address of the value this variable is pointing at
	// *pointer = give me the value this memory address is pointing at
	alexPointer := &alex
	alexPointer.updateName("Alexis")
	alex.print()

	// Of course theres a short cut for the above, golang smart enough to know that we are trying to manipulate the value
	// the pointer is referencing
	jim.updateName("Jimmy")
	jim.print()
}

// Receiver // This format of printing will print the field name
func (p person) print() {
	fmt.Printf("%+v", p)
}

// Important, golang is a pass by value language. The pointer issue. Go will make a copy of the value alex.updateName() but
// it will not be propagated back to the original value. To solve this, we need to use pointer.
func (pointerToPerson *person) updateName(newFirstName string) {
	// *pointer = give me the value this memory address is pointing at
	// IMPORTANT:
	// *person (above) type description: means we are working with a pointer that is pointing at a person type
	// *pointerToPerson (below) type description: means we are manipulating the value the pointer is referencing
	(*pointerToPerson).firstName = newFirstName
}
