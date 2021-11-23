package gotest

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := make(map[string]int)
	fmt.Printf("%#v\n", m)

	m["lym"] = 18
	m["hyy"] = 19
	m["wyw"] = 5

	fmt.Printf("%#v\n", m)

	for name, age := range m {
		fmt.Printf("%s\t%d\n", name, age)
	}

	m1 := map[string]int{}
	fmt.Printf("%#v\n", m1)
}

/*
内置的make函数可以创建一个map：
ages := make(map[string]int) // mapping from strings to ints

直接创建：
ages := map[string]int{
    "alice":   31,
    "charlie": 34,
}

这相当于

ages := make(map[string]int)
ages["alice"] = 31
ages["charlie"] = 34

因此，另一种创建空的map的表达式是map[string]int{}
*/

func TestCountChr(t *testing.T) {
	s := "hope is beautiful, it's for you"

	m2 := make(map[rune]int)

	for _, c := range s {
		m2[c]++
	}

	for chr, cnt := range m2 {
		fmt.Printf("%c\t%d\n", chr, cnt)
	}
}
