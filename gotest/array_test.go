package gotest

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestArray(t *testing.T) {
	a := [...]struct { // 结构体类型的数组
		name string
		age  int
	}{
		{"Bob", 10},
		{"Iris", 18},
		{"Lily", 15},
		{"Merry", 19}, // 逗号不能缺
	}

	fmt.Printf("%v\n", a)

	for i, e := range a {
		fmt.Println(i, e.name, e.age)
	}

	verify() // 验证golang中的数组是值传递的

	aa := [10]int{1, 5, 6}
	fmt.Println(sumArr(aa))

	// a1 := []int{1, 3, 5, 8, 7}
	// IndexValue1(a1, 8)
	// IndexValue2(a1, 8)

	a2 := []int{}
	fmt.Println(IndexValue1(GenSlice(a2), 100))
	// fmt.Println(IndexValue2(GenSlice(a2), 100))

}

func test(x [2]int) {
	fmt.Printf("x: %p\n", &x)
	x[1] = 1000
}

func verify() {
	x := [2]int{}
	fmt.Printf("x: %p\n", &x)
	// fmt.Println(x)
	// x[1] = 1000
	// fmt.Println(x)

	test(x)
	fmt.Println(x) // [0 0]
}

// 数组求和
func sumArr(n [10]int) int {
	sum := 0
	for i := 0; i < len(n); i++ {
		sum += n[i]
	}
	return sum
}

// 找出数组中和为给定值的两个元素的下标，例如数组[1,3,5,8,7]，
// 找出两个元素之和等于8的下标分别是（0，4）和（1，2）
func IndexValue1(a []int, target int) (cnt int) {
	for i := 0; i < len(a); i++ {
		other := target - a[i]
		for j := i + 1; j < len(a); j++ {
			if a[j] == other {
				fmt.Printf("[(%d, %d),(%d, %d)]\n", i, a[i], j, a[j])
				cnt++
				break
			}
		}
	}
	return
}

func IndexValue2(a []int, target int) (cnt int) {
	m := make(map[int]int)
	for i := 0; i < len(a); i++ {
		other := target - a[i]
		// 判断other有没有在map里， 如果有，打a[i],
		_, ok := m[other]
		if !ok {
			m[other] = i
		}
		j, ok := m[a[i]]
		if ok {
			fmt.Printf("[(%d, %d), (%d, %d)]\n", j, other, i, a[i])
			cnt++
		}
	}
	return
}

// 生成1000以内的随机整数并添加到slice
func GenSlice(a []int) []int {
	rand.Seed(time.Now().UnixNano()) // 生成种子数
	fmt.Println(rand.Intn(10))

	for i := 0; i < 100000; i++ {
		a = append(a, rand.Intn(1000))
	}
	return a
}
