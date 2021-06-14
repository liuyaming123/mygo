package main

import (
	"fmt"
	"gone/hello"
	// sli "gone/hello/myslice"
	// "gone/bubble"
	// "gone/countChr"
)

func main() {
	var p = fmt.Println
	p(hello.Hello())
	// fmt.Println(sli.Sli())

	// -------------------------------------------------------------------冒泡排序-------------------------------------------------------------------
	// "gone/bubble"
	// ss := []int{5, 5, 8, 3, 0, 6, 7, 1, 8, 3, 4}
	// fmt.Println(bubble.BubbleSort(ss))

	// -------------------------------------------------------------------统计字符-------------------------------------------------------------------
	// "gone/countChr"
	// fmt.Println(countChr.CntChr())

	// -------------------------------------------------------------------数组：数组为值类型-------------------------------------------------------------------
	// ar := [...]int{1, 2, 4} // ...代表根据元素确定数组长度，没有长度表示的是slice
	// fmt.Println(ar)
	// fmt.Printf("%T\n", ar)
	// ar1 := [5]int{} // 不指定值的话int默认为0，string默认为"", 也可以指定具体的索引值如{1: 99}
	// fmt.Println(ar1)

	// -------------------------------------------------------------------切片：切片为引用类型，因此在传递切片时将引用同一指针，修改值会影响其他的对象-------------------------------------------------------------------
	// s := []string{}                  // 初始化一个空slice
	// s = append(s, "apple", "banana") // 添加一个或多个元素
	// fmt.Println(s)
	// fmt.Printf("%T\n", s)
	// for i, it := range s { // slice的遍历
	// 	fmt.Println(i, it)
	// }

	// -------------------------------------------------------------------字符串-------------------------------------------------------------------
	// 给一个字符串后边添加一个一个字符串
	// st1 := "left"
	// st2 := "right"
	// st := st1 + ", " + st2
	// fmt.Println(st)

	// for i, r := range "hello, 世界" {
	// 	fmt.Printf("%d\t%c\t%d\n", i, r, r)
	// }

	// cnt := 0
	// for range "linux-mac" {
	// 	cnt++
	// }
	// fmt.Println(cnt)

	// p("Contains: ", s.Contains("test", "es"))
	// p("Count: ", s.Count("test", "t"))
	// p("HasPrefix: ", s.HasPrefix("test", "te"))
	// p("HasSuffix: ", s.HasSuffix("test", "st"))
	// p("ToLower: ", s.ToLower("TEst"))
	// p("ToUpper: ", s.ToUpper("TEst"))
	// p("char: ", "hello"[1])                         // 101
	// p("Index: ", s.Index("tesst", "s"))             // 第一个出现该字符的index
	// p("Join: ", s.Join([]string{"A", "B"}, "-"))    // 数组不能只用join操作
	// p("Replace: ", s.Replace("asss", "s", "y", -1)) // -1表示全部替换
	// p("Replace: ", s.Replace("assss", "s", "y", 2)) // 正整数表示依次替换几次

	// s1 := s.Split("a,b,c", ",")
	// p("Split: ", s1)
	// fmt.Printf("%T", s1) // slice类型

	// -------------------------------------------------------------------映射-------------------------------------------------------------------
	// m := map[string]int{}
	m := make(map[string]int, 100) // 指定容量的定义
	// m["a"] = 18  // 添加key-value
	// m["b"] = 19
	// p(m)
	// p(len(m))  // 取map的长度
	// delete(m, "a")  // 删除m中的key
	// p(m)
	// p(m["x"]) // 没有的值默认会取对应类型的默认值，string为"", int为0, bool为flase
	// p(m["a"])
	// m["a"] = 22 // 修改key-value
	// p(m["a"])

	// value, ok := m["c"] // 判断c是否在map中，如果存在，返回对应的value值，如果不存在，value值为对应类型的默认值，ok取值为true或者flase
	// p(value, ok)
	// if ok {
	// 	p(value)
	// } else {
	// 	p("not exist!")
	// }

	if value, ok := m["a"]; !ok { // 使用!ok 或者 ok来判断元素是否在map里
		p(value, ok)
	}

	// for key := range m { // 只遍历map的key
	// 	p(key, m[key])
	// }

	// for k, v := range m {  // k,v 同时遍历
	// 	p(k+":", v)
	// }

	// -------------------------------------------------------------------结构体-------------------------------------------------------------------
	type Employee struct {
		ID   int
		Name string
	}

	var yaming Employee
	yaming.ID = 1
	yaming.Name = "yaming"
	p(yaming)
}
