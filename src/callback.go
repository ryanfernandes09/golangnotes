package main

import (
	"fmt"
)

func callback() {

	randomNumbers := []int{43, 25, 653, 12, 9, 0}
	fmt.Println("Sum: ", sum(randomNumbers...))

	fmt.Println("Sum of even numbers: ", even(sum,randomNumbers...))

}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int{
	var evenNumbers []int
	for _, v := range vi {
		if v % 2 == 0 {
			evenNumbers = append(evenNumbers, v)
		}
	}
	return f(evenNumbers...)
}