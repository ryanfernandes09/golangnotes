package main

import (
	"fmt"
)

func canadianProvincesSlice() {

	slice := make([]string, 13, 13)
	slice = []string{"Newfoundland and Labrador",
		"Nova Scotia",
		"New Brunswick",
		"Prince Edward Island",
		"Quebec",
		"Ontario",
		"Manitoba",
		"Saskatchewan",
		"Alberta",
		"British Columbia",
		"Nunavut",
		"Northwest Territories",
		"Yukon"}

	fmt.Println("finalSlice: ", slice)
	fmt.Println("Length of the slice: ", len(slice))
	fmt.Println("Capacity of the slice: ", cap(slice))

	for i := 0; i < 13; i++ {
		fmt.Println(i, slice[i])
	}
}
