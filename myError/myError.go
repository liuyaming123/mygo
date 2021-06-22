package myError

import (
	"fmt"
	"os"
)

func TError(n int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered panic", err)
		}
	}()
	// defer func() {
	// 	fmt.Println("finally!")
	// }()
	if n < 10 {
		panic("n can't le 10") // panic退出前会调用defer，也会打印错误信息
	}
	fmt.Println("gte 10")
	os.Exit(1) // os.Exit退出时不会调用defer指定的函数， 也不输出当前调用栈信息

}
