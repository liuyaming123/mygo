package gotest

import (
	"fmt"
	"testing"
)

func TestMemory(t *testing.T) {
	// 示例代码
	// var a *int
	// *a = 100
	// fmt.Printf("a: %d\n", *a) // panic: runtime error: invalid memory address or nil pointer dereference

	// 示例代码中var a *int只是声明了一个指针变量a但是没有初始化，指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值。
	// 应该按照如下方式使用内置的new函数对a进行初始化之后就可以正常对其赋值了

	// var aa *int   // :1 声明类型的指针变量
	// aa = new(int) // :2 初始化为该类型的零值
	aa := new(int)   // 1, 2两步可合并为这一步
	fmt.Println(aa)  // 0xc000096090
	fmt.Println(*aa) // 0
	// *aa = 10
	fmt.Println(*aa) // 10

	b := make(map[string]int, 10)
	b["测试"] = 100
	fmt.Println(b)

	// make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，
	// 因为这三种类型就是引用类型，所以就没有必要返回他们的指针了

	f1() // 如果要从外面传入也可以使用引用类型的指针形式的变量，类似于python的给函数传list类型的参数

	mm := make(map[string]int, 5)
	fmt.Println(len(mm))
	for i := 0; i < 10; i++ {
		mm[fmt.Sprintf("%v", i)] = i
	}
	fmt.Println(len(mm))

	for k, v := range mm {
		fmt.Println(k, ": ", v)
	}
}

func f1() []int {
	res := make([]int, 0)
	res = append(res, 9)
	f2(&res)
	fmt.Println(res)
	return res
}

func f2(res *[]int) {
	*res = append(*res, 1, 4)
	fmt.Println(*res)
}
