package gotest

import (
	"fmt"
	"math"
	"testing"
)

func TestString(t *testing.T) {
	s := "pprof.cn博客"
	for i := 0; i < len(s); i++ { // byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, c := range s { // rune
		fmt.Printf("%v(%c) ", c, c)
	}
	fmt.Println()
	// 因为UTF8编码下一个中文汉字由3~4个字节组成，所以我们不能简单的按照字节去遍历一个包含中文的字符串，
	// 否则就会出现上面输出中第一行的结果。

	//字符串底层是一个byte数组，所以可以和[]byte类型相互转换。
	// 字符串是不能修改的 字符串是由byte字节组成，所以字符串的长度是byte字节的长度。
	// rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。

	fmt.Printf("字节编码: %#v\n", []byte("pr博客"))
	fmt.Printf("utf-8编码: %#v\n", []rune("pr博客"))
}

func TestSqrtDemo(t *testing.T) {
	// 类型转换
	var a, b = 3, 4
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c := int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func TestModifyString(t *testing.T) {
	// 要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
	s := "hello"
	sByte := []byte(s)
	sByte[0] = 'H'
	sNew := string(sByte)
	fmt.Println(sNew)
	fmt.Println(&s == &sNew) // false
}
