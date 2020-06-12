package main

import (
	"fmt"
)

func playground() {
	fmt.Println(foo3())
	fmt.Println(bar3())

}

func foo3() int {
	return 5
}

func bar3() (int, string) {
	return 6, "six"
}
