package structs

import "math"

type Shape interface {
	Area() float64
}
type Rectangle struct {
	width  float64
	height float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.width * rectangle.height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * math.Pi * c.Radius
}

func Perimeter(rectangle Rectangle) float64 {
	return (2 * (rectangle.width + rectangle.height))
}
