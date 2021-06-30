package gotest

import (
	"fmt"
	"strconv"
	"testing"
)

func TestStr(t *testing.T) {
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1) //type:int value:100
	}

	fmt.Println(strtoint(s1))
}

func strtoint(s string) (i int, err error) {
	i, err = strconv.Atoi(s)
	return
}

func TestByte(t *testing.T) {
	var aa byte = 'a'       // 单个字符 byte(uint8类型的别名)。 可以和整型类型直接相加减
	fmt.Printf("%c \n", aa) // a 占位符%c表示字符
	fmt.Printf("%T \n", aa) // uint8 (byte类型其实就是uint8类型的别名)
}
