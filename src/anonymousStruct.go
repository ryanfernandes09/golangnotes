package main

import (
	"fmt"
)

func anonymousStruct() {

	s1 := struct {
		var1 map[int]string
		var2 []string
		var3 string
	}{
		var1: map[int]string{
			1: "first",
			2: "second",
		},
		var2: []string{"apples", "and", "oranges"},
		var3: "idioms",
	}

	fmt.Println("map value: ", s1.var1)
	fmt.Println("slice value: ", s1.var2)
	fmt.Println("string value: ", s1.var3)
	for k, v := range s1.var1 {
		fmt.Printf("value at position %d : %s \n", k, v)
	}

}
