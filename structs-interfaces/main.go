package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.width + r.height)
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	rad float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.rad * c.rad
}

type Shape interface {
	Area() float64
}

func main() {
	fmt.Println("tatft")
}
