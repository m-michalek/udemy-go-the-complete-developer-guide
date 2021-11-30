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
}
