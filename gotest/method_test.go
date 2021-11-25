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
/*
func (p *Point) Distance(q Point){}  && func (p Point) Distance(q Point) float64 {}
不管你的method的receiver是指针类型还是非指针类型，都是可以通过指针/非指针类型进行调用的，编译器会帮你做类型转换。
在声明一个method的receiver该是指针还是非指针类型时，你需要考虑两方面的因素，第一方面是这个对象本身是不是特别大，如果声明为非指针变量时，调用会产生一次拷贝；第二方面是如果你用指针类型作为receiver，那么你一定要注意，这种指针类型指向的始终是一块内存地址，就算你对其进行了拷贝。熟悉C或者C++的人这里应该很快能明白。
*/
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
