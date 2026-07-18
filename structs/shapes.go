package main

import "math"

// Shape is implemented by anything that can tell us its Area
type Shape interface {
	Area() float64
}

// Rectangle has the dimensions of a rectangle.
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter returns the perimeter of a rectangle.
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Circle represents a circle...
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
