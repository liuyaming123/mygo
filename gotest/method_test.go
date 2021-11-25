package gotest

import (
	"fmt"
	"math"
	"testing"
)

type Point struct {
	X, Y float64
}

type Path []Point

/*
我们可以给同一个包内的任意命名类型定义方法，只要这个命名类型的底层类型
（译注：这个例子里，底层类型是指[]Point这个slice，Path就是命名类型）
不是指针或者interface
*/

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
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

	p3 := Path{{0, 0}, {0, 1}, {0, 5}, {0, 3}}
	fmt.Println(p3.Distance())
}
