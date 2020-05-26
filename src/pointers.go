package main

import (
	"fmt"
)

type person struct {
	address string
}

func changeme(p1 *person) {
	fmt.Println("p1 : ", &p1)
	p1.address = "new addr" // (*p1).address and p1.address has the same effect
	//(*p1).address = "new address"
}

func pointers() {
	x := 5
	fmt.Println("address of x : ", &x)
	fmt.Println("value of x: ", x)

	pMain := person{
		address: "old address",
	}
	fmt.Println(pMain.address)
	changeme(&pMain)
	fmt.Println(pMain.address)
}
