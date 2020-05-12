package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("\tMENU")
	fmt.Println("0. playground")
	fmt.Println("1. arraySlice")
	fmt.Println("2. weirdMaps")
	fmt.Println("3. canadianProvincesSlice")
	fmt.Println("4. simpleMap")
	fmt.Println("5. structs")
	fmt.Println("6. Exit")
	fmt.Printf("Enter your choice: ")

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	switch {
	case input.Text() == "0":
		playground()
	case input.Text() == "1":
		arraySlice()
	case input.Text() == "2":
		weirdMaps()
	case input.Text() == "3":
		canadianProvincesSlice()
	case input.Text() == "4":
		simpleMap()
	case input.Text() == "5":
		structs()
	case input.Text() == "6":
		os.Exit(0)
	default:
		fmt.Println("Please enter a valid choice.")
	}
}
