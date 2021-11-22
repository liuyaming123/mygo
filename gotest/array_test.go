package gotest

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestArray(t *testing.T) {
	a := [...]struct { // 元素是结构体类型的数组
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

	// a2 := []int{}
	// fmt.Println(IndexValue1(GenSlice(a2), 100))
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

func TestArrSli(t *testing.T) {
	array := [4]int{10, 20, 30, 40}
	slice := array[0:2]
	newSlice := append(slice, 50)
	fmt.Printf("one slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("one newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	newSlice[1] += 10
	fmt.Printf("two slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("two newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
}

func TestMonth(t *testing.T) {
	months := [...]string{1: "Jan", 2: "Feb", 3: "Mar", 4: "Apr", 5: "May", 6: "Jun", 7: "Jul", 8: "Aug", 9: "Sep", 10: "Oct", 11: "Nov", 12: "Dec"}
	for i, m := range months {
		fmt.Println(i, m)
	}

	Q2 := months[4:7]
	fmt.Println(Q2, len(Q2), cap(Q2)) // [Apr May Jun] 3 9

	// 如果切片操作超出cap(s)的上限将导致一个panic异常，但是超出len(s)则是意味着扩展了slice，因为新slice的长度会变大
	fmt.Println(Q2[:5]) // 输出: [Apr May Jun Jul Aug];  explain: extend a slice (within capacity)
	// fmt.Println(Q2[:10]) // panic: out of range

	/*
		字符串的切片操作和[]byte字节类型切片的切片操作是类似的。都写作x[m:n]，并且都是返回一个原始字节序列的子序列，底层都是共享之前的底层数组，因此这种操作都是常量时间复杂度。x[m:n]切片操作对于字符串则生成一个新字符串，如果x是[]byte的话则生成一个新的[]byte。
	*/
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func TestReverse(t *testing.T) {
	a := []int{1, 3, 5, 7, 9}
	reverse(a)
	fmt.Println(a)

	// slice动态扩容机制
	fmt.Println(a)
	fmt.Println(cap(a)) // 5  首次容量为初始化的slice长度
	a = append(a, 2, 3, 4, 4, 4, 6, 2, 4)
	fmt.Println(a)
	fmt.Println(cap(a)) // 为了提高内存使用效率，append超过原容量时，新数组分配的容量会略大于数组长度
}
