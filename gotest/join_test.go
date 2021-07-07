package gotest

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestJoin(t *testing.T) {
	num := 50000
	list := make([]string, 10)
	for i := 0; i < num; i++ {
		list = append(list, "helloWorld")
	}

	var test1, seq string
	seq = " "
	start1 := time.Now()
	for i := 0; i < num; i++ {
		test1 += list[i] + seq
	}
	end1 := time.Since(start1)

	fmt.Println(end1)

	start2 := time.Now()
	test2 := strings.Join(list, " ")
	end2 := time.Since(start2)
	println(test2[:100])
	println(end2)
}
