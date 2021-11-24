package gotest

import (
	"fmt"
	"math"
	"testing"
)

type Point struct {
	X, Y float64
}

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, bug as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func TestMethod(t *testing.T) {
	p1 := Point{0, 0}
	p2 := Point{4, 5}

	fmt.Printf("%#v\n", p1.Distance(p2))
}
