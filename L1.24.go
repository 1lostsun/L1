package main

import (
	"fmt"
	"math"
)

func main() {
	point1 := NewPoint(2.4, 3.5)
	point2 := NewPoint(12.9, 5.78)

	fmt.Println(point1.DistanceTo(point2))
}

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p *Point) DistanceTo(other *Point) float64 {
	xDist := math.Abs(p.x - other.x)
	yDist := math.Abs(p.y - other.y)
	return math.Sqrt(xDist*xDist + yDist*yDist)
}
