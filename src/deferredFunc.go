package main

import (
	"fmt"
)

func deferredFunc() {
	fmt.Println("\n <-------deferredFunc.go--------> ")

	fmt.Println("this is useful for closing files since the order of writing code \ncan be open-file --> close-file --> file operations \nbut the order of execution will still be  \nopen-file --> file-operations --> close-file\n\n")

	fmt.Println("Opening a file")
	defer deferred()
	fmt.Println("File operations...")

}

func deferred() {
	fmt.Println("closing files...")
}
