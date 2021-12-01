package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

type business struct {
	name        string
	contactInfo contactInfo
}

type contactInfo struct {
	mail string
	zip  string
}

// if the name of the property and the type are equal it can be aggregated
type school struct {
	name string
	contactInfo
}

func (p person) print() {
	fmt.Printf(p.firstName+": %+v\n", p)
}

func (p person) updateNameIncorrectly(name string) {
	// Go passes variables by value which mean that p person is a copy of what is actually passed to the function
	// manipulting the copy will not manipulate the actual construct where the function was called
	// here you can use pointers like shown in updateName()
	p.firstName = name
}

// *person means that the function expects a pointer to a person
// *pointerOfPerson here the * is an actual operator which gives access to the value (the actual struct in memory) instead of the address in memory
func (pointerOfPerson *person) updateName(name string) {
	(*pointerOfPerson).firstName = name
}

func main() {
	fmt.Println("1. approach to construct a struct")
	peter := person{"Peter", "Doe"}
	fmt.Printf("peter: %v\n", peter)

	fmt.Println("2. approach to construct a struct")
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Printf("alex: %v\n", alex)

	fmt.Println("3. approach to construct a struct")
	var mike person
	fmt.Printf("mike: %v\n", mike)
	fmt.Printf("mike: %+v\n", mike)
	mike.firstName = "Mike"
	mike.lastName = "Morris"
	fmt.Printf("mike: %v\n", mike)
	fmt.Printf("mike: %+v\n", mike)

	fmt.Println()

	fmt.Println("Embedded/nested structs")
	microsoft := business{
		name: "Microsoft",
		contactInfo: contactInfo{
			mail: "info@microsoft.com",
			zip:  "123",
		},
	}
	fmt.Printf("microsoft: %+v\n", microsoft)

	hwrBerlin := school{
		name: "Berlin School of Economics and Law Berlin",
		contactInfo: contactInfo{
			mail: "info@htw-berlin.com",
			zip:  "123",
		},
	}
	fmt.Printf("hwrBerlin: %+v\n", hwrBerlin)

	fmt.Println("Adding custom functions to structs")
	peter.updateNameIncorrectly("Pet")
	peter.print()

	peterPointer := &peter
	peterPointer.updateName("Ron")
	peterPointer.print()

	// it is not necessary to explicitly pass the pointer to updateName() to change the value
	peter.updateName("Pet")
	peter.print()
}
