package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
}

func main() {
	// ac := person{"Ananta", "Chakravartula"}
	// fmt.Println(ac)
	// recommended way to make sure changing order of struct definition will not mess up defined properties
	sc := person{firstName: "Shilpa", lastName: "Chakravartula"}
	// go doesn't have nil/none, it will add zero values
	ac := person{
		firstName: "Ananta",
		lastName:  "Chakravartula",
		//contact:   contactInfo{email: "ac@foo.com", zipCode: "01810"},
		contactInfo: contactInfo{email: "ac@foo.com", zipCode: "01810"},
	}
	fmt.Println(sc)
	fmt.Printf("%+v \n", sc)
	fmt.Printf("%+v \n", ac)
	ac.print()
	ac.updateName("Foo")
	// oops First name will be still Ananta :). Go makes a copy and updates
	ac.print()

	acPointer := &ac
	acPointer.newName("Foo")
	ac.print()
	// Go shortcut. It will automatically convert person struct to pointer
	ac.newName("Qux")
	ac.print()

}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (pointerToPerson *person) newName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("Printing from receiver function: %+v \n", p)
}
