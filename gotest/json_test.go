package gotest

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Student struct {
	Id     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
}

type Class struct {
	Title    string     `json:"title"`
	Students []*Student `json:"students"`
}

func TestJson(t *testing.T) {
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			Id:     i,
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
		}
		c.Students = append(c.Students, stu)
	}

	// json序列化：结构体 ---> json格式的字符串
	data, err := json.MarshalIndent(c, " ", "\t")
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json: %s\n", data)

	// json反序列化： json格式的字符串 ---> 结构体
	str := `{"Title":"101","Students":[{"Id":0,"Gender":"男","Name":"stu00"},{"Id":1,"Gender":"男","Name":"stu01"},{"Id":2,"Gender":"男","Name":"stu02"},{"Id":3,"Gender":"男","Name":"stu03"},{"Id":4,"Gender":"男","Name":"stu04"},{"Id":5,"Gender":"男","Name":"stu05"},{"Id":6,"Gender":"男","Name":"stu06"},{"Id":7,"Gender":"男","Name":"stu07"},{"Id":8,"Gender":"男","Name":"stu08"},{"Id":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json Unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
	for i, student := range c1.Students {
		fmt.Println(i, fmt.Sprintf("%#v", *student))
	}
}

//Student 学生
type Student_ struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

func TestTag(t *testing.T) {
	s1 := Student_{
		ID:     1,
		Gender: "女",
		name:   "pprof",
	}
	data, err := json.MarshalIndent(s1, " ", "\t")
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:\n%s\n", data) //json str:{"id":1,"Gender":"女"}
}
