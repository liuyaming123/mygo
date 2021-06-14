package myset

import (
	"fmt"
	"testing"
)

func TestMySet(t *testing.T) {
	// m := map[int]func(op int) int{}
	// m[1] = func(op int) int { return op }
	// m[2] = func(op int) int { return op * op }
	// m[3] = func(op int) int { return op * op * op }
	// t.Log(m[1](2), m[2](2), m[3](2))
	// print("hehe\n")

	mySet := map[int]bool{}
	mySet[1] = true
	mySet[2] = true
	if mySet[1] {
		fmt.Printf("%d is exist!\n", 1)
	} else {
		fmt.Printf("%d is not exist\n", 1)
	}

}
