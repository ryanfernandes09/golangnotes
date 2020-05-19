package main

import (
	"fmt"
)

func callback() {

	randomNumbers := []int{43, 25, 653, 12, 9, 0}
	s1 := sum(randomNumbers...)

	fmt.Println("Sum: ", s1)

}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
