package main

import (
	"fmt"
	"math"
)

type square struct{
	color string
	name string
	lengthOfSide float64
}

type circle struct{
	color string
	name string
	radius float64
}

type shape interface{
	area() float64
}

func (s square) area() float64{
	area := s.lengthOfSide*s.lengthOfSide
	return area
}

func (c circle) area() float64{
	area := math.Pi * c.radius * c.radius
	return area
}

func info(s shape) {
	fmt.Printf("\nArea of a %T which is received by function area: %f\n",s,s.area())
}

func inheritance() {

	s := square{
		color: "black",
		name: "square",
		lengthOfSide: 56,
	}

	c := circle{
		color: "blue",
		name: "cirle",
		radius: 23,
	}

	info(s)
	info(c)
	fmt.Println("The function area is shared")

}
