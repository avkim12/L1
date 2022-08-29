package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func main() {

	var a, b, c float64

	p1 := NewPoint(-1, 2)
	p2 := NewPoint(-3, 4)

	a = math.Abs(p1.x) + math.Abs(p2.x)
	b = math.Abs(p1.y) + math.Abs(p2.y)
	c = math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	
	fmt.Printf("%.10f", c)
}