package main

import (
	"fmt"
)

//this is a good example to understand that functions
//are types in go.
func closure() {

	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())

}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
