package main

import (
	"fmt"
	"math"
)

// Define the shape interface with an area method.
type shape interface {
	area() float64
}

// Define the circle struct with a radius field.
type circle struct {
	radius float64
}

// Define the rectangle struct with width and height fields.
type rectangle struct {
	width  float64
	height float64
}

// Implement the area method for circle type using a method receiver.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Implement the area method for rectangle type using a method receiver.
func (r rectangle) area() float64 {
	return r.width * r.height
}

func main() {
	fmt.Println("Interface in Go")

	// Create instances of circle and rectangle with the required fields.
	c1 := circle{radius: 4.5}
	r1 := rectangle{width: 4, height: 5}

	// Create a slice of shape interface, holding both circle and rectangle instances.
	shapes := []shape{c1, r1}

	// Iterate through the slice and print the area of each shape.
	for _, s := range shapes {
		fmt.Println(s.area())
	}
}
