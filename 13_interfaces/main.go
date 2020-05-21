package main

import (
	"fmt"
	"math"
)

// Interfaces are list of methods that you can implement on different struct

//Shape Interface
type Shape interface {
	area() float64
}

//Circle Struct
type Circle struct {
	x, y, radius float64
}

// Rectangle Struct
type Rectangle struct {
	l, b float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.l * r.b
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{0, 0, 12}
	rectangle := Rectangle{5, 10}

	fmt.Printf("Circle - %f\n", getArea(circle))
	fmt.Printf("Rectangle - %f\n", getArea(rectangle))
}
