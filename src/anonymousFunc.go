package main

import(
	"fmt"
)

func anonymousFunc(){
	foo()

	func(){
		fmt.Println("anonymous func ran")
	}()

	func(x int){
		fmt.Println("THe meaning of life:",x)
	}(42)
}

func foo(){

	fmt.Println("Foo ran")
}