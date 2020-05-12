package main

import (
	"fmt"
)

func playground() {

	p1 := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "james",
		lastName:  "bond",
		age:       42,
	}

	fmt.Println("p1: ", p1)
}
