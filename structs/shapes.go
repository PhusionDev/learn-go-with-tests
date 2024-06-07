package structs

import "math"

// structs
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// interfaces
type Shape interface {
	Area() float64
}

// Rectangle methods
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle methods
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle methods
func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
