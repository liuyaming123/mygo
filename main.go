package main

import (
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	// "gone/feibo"
	// "gone/hello"
	// "gone/myError"
	// "gone/goRoutinePool"

	// "gone/myInterface"  // 多态
	"unsafe"
	// sli "gone/hello/myslice"
	// "gone/bubble"
	// "gone/countChr"
)

var p = fmt.Println
var pf = fmt.Printf

func main() {
	// goRoutinePool.MyJob()
	// CC5()
	// CC4()
	// CC3()
	// CC2() // 主协程结束，所有的守护协程也就结束
	// CC1()
	// myError.TError(8)
	// p(feibo.FeiBo(10))
	// p(hello.Hello())
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
	// m := make(map[string]int, 100) // 指定容量的定义
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

	// if value, ok := m["a"]; !ok { // 使用!ok 或者 ok来判断元素是否在map里
	// 	p(value, ok)
	// }

	// for key := range m { // 只遍历map的key
	// 	p(key, m[key])
	// }

	// for k, v := range m {  // k,v 同时遍历
	// 	p(k+":", v)
	// }

	// -------------------------------------------------------------------机构体-------------------------------------------------------------------
	// Struct()

	// -------------------------------------------------------------------defer-------------------------------------------------------------------
	// p(TDefer())

	// -------------------------------------------------------------------接口-------------------------------------------------------------------
	// o := &stockPosition{"Nike", 1.20, 4}
	// showValue(o)
	// p(o.getValue())
	// var t f
	// showValue(t)
	// p(t.getValue())

	// -------------------------------------------------------------------扩展与复用-------------------------------------------------------------------
	// MyDog()
	// -------------------------------------------------------------------多态-------------------------------------------------------------------
	// myInterface.MyIf()

	// MyGroutine()
	// WriteFile()
}

func Clear() {
	fmt.Println("clear resources.")
}

func TDefer() string {
	defer Clear() // 会在return之前执行，即使有panic，也会执行
	fmt.Println("Start...")
	return "jaja"
}

type Employee struct {
	ID   int
	Name string
	Age  int
}

func Struct() {
	// 初始化方式
	e := Employee{0, "Bob", 10}
	// e := Employee{Name: "Bob", Age: 20}
	p(e)

	e2 := new(Employee) // 注意这里返回的引用（指针），相当于e := &Employee{}
	e2.Age = 100
	e2.ID = 2
	e2.Name = "lily"
	p(e2)
	pf("e type is: %T\n", e)
	pf("e2 type is: %T\n", e2)
	p("addr is %x", unsafe.Pointer(&e2.Name))
	// p(&e2.Name)
	p(e2.String())
}

func (e *Employee) String() string { //定义结构体方法，使用指针只用一份数据，推荐使用指针
	p("addr is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("Id: %d - Name: %s - Age: %d", e.ID, e.Name, e.Age)
}

// func (e Employee) String() string { //定义结构体方法，使用原始类型会copy出一份数据， 不推荐
// 	p("addr is %x", unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("Id: %d - Name: %s - Age: %d", e.ID, e.Name, e.Age)
// }

// 变长参数的函数
func Sum(vals ...int) int {
	fmt.Println(vals)
	fmt.Printf("vals type is %T\n", vals) //[]int
	total := 0
	for _, v := range vals {
		total += v
	}
	return total
}

func IsTrue(name string) (yes1 bool, yes2 bool, err error) {
	if len(name) > 1 {
		yes1 = true
	}
	return
}

// 接口：是用来定义对象之间交互的协议的。
type stockPosition struct {
	ticket     string
	sharePrice float32
	count      float32
}

func (s *stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type valueable interface {
	getValue() float32
}

func showValue(asset valueable) {
	fmt.Printf("value of the asset is %f\n", asset.getValue())
}

type f float32

func (s f) getValue() float32 {
	return 0.3
}

// 扩展与重用
type Pet struct {
}

func (pet *Pet) Speak() {
	p("...")
}

func (pet *Pet) SpeakTo(host string) {
	pet.Speak()
	p(" ", host)
}

type Dog struct {
	Pet
}

func MyDog() {
	d := new(Dog)
	d.SpeakTo("hyy")
}

func MyGroutine() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			p(i)
		}(i)
		time.Sleep(time.Microsecond)
	}
}

func WriteFile() {
	for i := 0; i < 10; i++ {
		WriteWithIoutil(fmt.Sprintf("./t_%d.txt", i), fmt.Sprintf("中华上下五千年%d\n", i))
	}
}

func WriteWithIoutil(name, content string) {
	defer wg.Done() // goroutine结束就登记-1
	data := []byte(content)
	if ioutil.WriteFile(name, data, 0644) == nil {
		fmt.Println("写入文件成功:", content)
	}
}

var wg sync.WaitGroup

func Hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
	// time.Sleep(time.Microsecond)
}

func CC1() {
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1) // 启动一个goroutine就登记+1
	// 	go Hello(i)
	// }
	// wg.Wait() // 等待所有登记的goroutine都结束

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go WriteWithIoutil(fmt.Sprintf("./t_%d.txt", i), fmt.Sprintf("中华上下五千年%d\n", i))
	}
	wg.Wait()
}

func CC2() {
	// 合起来写
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i = %d\n", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		time.Sleep(time.Second * 5)
		fmt.Printf("main goroutine: i = %d\n", i)
		if i == 2 {
			break
		}
	}
}

func recv(c chan int) {
	defer wg.Done()
	ret := <-c
	p("接收成功, 接收值是:", ret)
}

func sen(c chan int) {
	defer wg.Done()
	c <- 'A'
	p("发送成功")
}

func CC3() {
	// var ch chan int
	// p(ch) // 通道类型的控制是nil, 声明的通道后需要使用make函数初始化之后才能使用

	// ch1 := make(chan int)
	// wg.Add(2)
	// go recv(ch1)
	// go sen(ch1)
	// wg.Wait()
	// close(ch1)

	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c) // 当通道被关闭时，往该通道发送值会引发panic，从该通道里接收的值一直都是类型零值
	}()

	for {
		if data, ok := <-c; ok {
			p(data)
		} else {
			break
		}
	}

	p("Over")

}

func CC4() {
	ch1 := make(chan int) //无缓冲通道
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()

	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		p(i)
	}
	p(ch2) // 0xc0000820c0
}

func counter(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func CC5() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}
