package mytest

import (
	"fmt"
	"testing"
)

func modify1(x int) {
	x += x
}

func modify2(x *int) {
	*x = 100
}

func TestPoint(t *testing.T) {
	// a := 10
	// b := &a
	// fmt.Printf("a: %d, ptr: %p\n", a, &a)
	// fmt.Printf("b: %p, type: %T\n", b, b)
	// fmt.Println(&b)

	a := 10
	modify1(a)
	fmt.Println(a)
	modify2(&a)
	fmt.Println(a)

	var s *int     // 指针声明
	fmt.Println(s) //  <nil> 空指针;
	// var a *int只是声明了一个指针变量a但是没有初始化，指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值。
	s = new(int)    // 指针初始化
	fmt.Println(*s) // 初始化之后就是int的零值0

	// 上面两步可以合并为一步：
	ss := new(bool) // 声明加初始化
	fmt.Println(*ss)

	m := new(map[string]int)
	fmt.Println(*m) // map[]
	sss := new([]int)
	fmt.Println(*sss) // []

	// make也是用于分配内存的，区别于new，它只用于slice、map以及chan类型的内存创建，而且它返回的类型就是这三个类型本身，
	//而不是它们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回它们的指针了。
	// make函数是无可替代的，我们在使用slice、map以及chan的时候，都需要使用make进行初始化。
	b := make(map[string]int, 10)
	fmt.Println("b:", b) // b: map[]
	b["age"] = 18
	fmt.Println("b:", b)

	var aa int
	fmt.Println(&aa, aa)
	p := &aa
	*p = 20
	fmt.Println(&aa, aa) // 指针指向的值变了，原本的值就变了
}
