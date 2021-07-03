package gotest

import (
	"fmt"
	"testing"
)

type student struct {
	name string
	age  int
}

func TestStruct(t *testing.T) {
	students := []student{{"wjc", 33}, {"kelvin", 11}}
	stuMap := make(map[int]*student)
	for i, stu := range students {
		stuMap[i] = &stu
		fmt.Printf("1----> %p,%p,%+v\n", &i, &stu, stu)
		// 循环过程中，stu变量只声明了一次，所以stu地址即&stu是不变的，值是变化的。所以&stu始终不变
		// 同理，变量i也只声明了一次，所以i的地址即&i也是不变的
	}

	for k, v := range stuMap {
		fmt.Println("2----> ", k, v)
		// 此时，值都变成了 kelvin
	}

	for i := range students { // 只遍历索引
		stuMap[i] = &students[i]
		fmt.Printf("3----> %p,%+v\n", &students[i], students[i])
	}

	for k, v := range stuMap {
		fmt.Println("4----> ", k, v)
	}
}

func TestStruct2(t *testing.T) {
	s1 := newStudent("hyy", 17)
	fmt.Printf("5----> %#v\n", s1)

}

// 因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，
// 下面是一个返回结构体指针类型的构造函数
func newStudent(name string, age int) *student {
	return &student{
		name: name,
		age:  age,
	}
}
