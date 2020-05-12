package main

import (
	"fmt"
)

func simpleMap() {
	user_data := map[string][]string{
		"bond_james":      []string{"shaken, not stirred", "martinis", "women"},
		"moneypenny_miss": []string{"james bond", "literature", "computer science"},
		"no_dr":           []string{"being evil", "ice cream", "sunsets"},
	}

	fmt.Println(user_data)

	delete(user_data, "bond_james")

	for k, v := range user_data {
		fmt.Println("For ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

}
