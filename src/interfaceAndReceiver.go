package main

import(
	"fmt"
)

func (s secretAgent) speak(){
	fmt.Println(s.firstName," can speak.")
	fmt.Println("The names ",s.lastName,", James Bond.")
}

type human interface{
	speak()
}

type secretAgent struct {
	firstName string
	lastName string
	ltk bool
}

func interfaceAndReceiver(){
	sa1 := secretAgent{
		firstName: "James",
		lastName: "Bond",
		ltk: true,
	}

	sa1.speak()
}