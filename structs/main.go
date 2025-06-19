package main

import "fmt"

type ContactInfo struct {
	email   string
	zipCode int
}

type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

func main() {
	// var alex Person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	jim := Person{
		firstName: "Jim",
		lastName:  "Party",
		contact: ContactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}

	// alex.print()
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *Person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}
