package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func PrintArea(x shape) {
	fmt.Printf("%T\n", x)
	fmt.Println("Area of the shape is - ", x.area())
}

func main() {
	s1 := square{10}
	c1 := circle{5}
	PrintArea(s1)
	PrintArea(c1)
}
