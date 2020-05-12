package main

import (
	"fmt"
)

func weirdMaps() {
	fmt.Println("\n <-------weirdMaps.go--------> ")

	m := map[string]int{
		"Ryan":   26,
		"Aurene": 23,
	}

	fmt.Println("\nPrint the whole map here...", m)

	fmt.Println("\nTrying to retrieve a key that does not exist in the map returns 0 and not a null.")
	fmt.Println("the expression `if v, ok := m[\"MissingKey\"]: ok {... \nis commonly used to work on map retrievals.")

	fmt.Println("\nretrieving age of Ryan: ", m["Ryan"])
	fmt.Println("retrieving age of Aurene: ", m["Aurene"])
	fmt.Println("retrieving age of MissingKey: ", m["MissingKey"])

	fmt.Println("\nDeleting from a map is done by using the delete function -->  delete(map name, key name)")
	delete(m, "Ryan")
	fmt.Println("After deleting Ryan from the map, the map is : ", m)
}
