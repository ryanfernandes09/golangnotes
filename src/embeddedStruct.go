package main

import (
	"fmt"
)

func embeddedStruct() {
	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	t1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "orange",
		},
		fourWheel: false,
	}

	t2 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: true,
	}

	fmt.Println("Truck: ", t1)
	fmt.Println("Sedan: ", t2)

	fmt.Printf("\nThe truck has %d doors", t1.doors)
}
