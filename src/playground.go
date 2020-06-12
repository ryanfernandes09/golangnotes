package main

import (
	"fmt"
)

func playground() {

	x := returner()
	x()

}

func returner() func() {
	return printer
}

func printer() {
	fmt.Println("in the printer")
}
