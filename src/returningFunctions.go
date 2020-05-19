package main

import(
	"fmt"
)

/*
Avoid functional programming practices like this as much as possible
for code thats easier to understand.
*/

func returningFuncs(){

	fmt.Println(foo()())

	func foo() func{
		return bar()
	}
	
	func bar() int{
		return 42
	}
}