package main

import "fmt"

func arraySlice() {
	fmt.Println("\n <-------arraySlice.go--------> ")

	/***Every slice references an underlying array. The program below shows
	how the underlying array expands to accomodate new values in a slice.
	memory management is key when there are underlying expansions that the
	developer is not always aware about.
	***/
	x := make([]int, 10, 12)

	fmt.Println("\nthis shows how the array in the background expands as the slice grows. \nslices are stored as arrays in the background")

	fmt.Print("X is: ", x)
	fmt.Print("\t\tlength of x: ", len(x))
	fmt.Println("\tcapacity of x: ", cap(x))

	x = append(x, 1)
	fmt.Print("X is: ", x)
	fmt.Print("\t\tlength of x: ", len(x))
	fmt.Println("\tcapacity of x: ", cap(x))

	x = append(x, 2)
	fmt.Print("X is: ", x)
	fmt.Print("\t\tlength of x: ", len(x))
	fmt.Println("\tcapacity of x: ", cap(x))

	x = append(x, 3)
	fmt.Print("X is: ", x)
	fmt.Print("\tlength of x: ", len(x))
	fmt.Println("\tcapacity of x: ", cap(x))

	fmt.Println("\nMulti-dimensional slices are a slice <of a slice of integers/strings/etc>")
	z := []int{1, 2, 3}
	y := [][]int{x, z}

	fmt.Println(y)
}
