package main

import(
	"fmt"
)

/*** here, the function speak() has a receiver. This means that every value of type 
receiver which is in this case, secretAgent will receive the function speak().

Furthermore, the interface human has the function speak(). This implies that every
value that has the function speak() is of type human.

Summary : Value sa1 has function speak by virtue of being of type secretAgent
It is also of type human because it has the function speak.
***/ 

func (s secretAgent) speak(){
	fmt.Println(s.firstName," can speak.")
	fmt.Println("The names ",s.lastName,", James Bond.")
}

type human interface{
	speak()
}

type secretAgent struct {
	firstName string
	lastName string
	ltk bool
}

func interfaceAndReceiver(){
	sa1 := secretAgent{
		firstName: "James",
		lastName: "Bond",
		ltk: true,
	}

	sa1.speak()
	fmt.Println(human(sa1))

}