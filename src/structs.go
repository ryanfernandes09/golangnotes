package main

import(
	"fmt"
)

func structs(){

	type person struct{
		firstName string
		lastName string
		favFlavours []string
	}

	ryan := person{
		firstName: "Ryan",
		lastName: "Fernandes",
		favFlavours: []string{"butterscotch", "strawberry", "tiramisu"},
	}

	bond := person{
		firstName: "James",
		lastName: "Bond",
		favFlavours: []string{"aston", "martin", "chocolate"},
	}

	fmt.Println("\n\nCreated a struct with attributes firstName (string) lastName (string) and favFlavours (slice[string])\n")
	fmt.Println("ryan: ",ryan)
	fmt.Println("bond: ",bond)
}