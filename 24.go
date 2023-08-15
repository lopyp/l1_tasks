package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func distance(p1, p2 Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := Point{x: 0, y: 0}
	p2 := Point{x: 3, y: 4}
	d := distance(p1, p2)
	fmt.Println("Distance between", p1, "and", p2, "is", d)
}
