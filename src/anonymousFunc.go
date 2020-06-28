package main

import (
	"fmt"
)

func anonymousFunc() {
	foo()

	func() {
		fmt.Println("anonymous func ran")
		fmt.Println("test vscode git plugin")
	}()

	func(x int) {
		fmt.Println("THe meaning of life:", x)
	}(42)
}

func foo() {

	fmt.Println("Foo ran")
}
