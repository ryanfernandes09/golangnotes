package main

import (
	"fmt"
)

/*
Avoid functional programming practices like this as much as possible
for code thats easier to understand.
*/

func returningFunctions() {
	fmt.Println(secondaryFunc()())
}

func primaryFunc() int {
	return 451
}

func secondaryFunc() func() int {
	return func() int {
		return primaryFunc()
	}
}
