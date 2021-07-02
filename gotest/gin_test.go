package gotest

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGin(t *testing.T) {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		// var msg []string //定义变量
		// type msg2 []string //定义别名
		var msg struct {
			Name string `json:"name3"`
			Age  int
			Num  []int `json:"number"`
		}
		msg.Name = "hyy"
		msg.Age = 18
		msg.Num = []int{1, 3, 4}
		c.JSON(200, msg)
	})
	r.Run(":8888") // listen and serve on 0.0.0.0:8080
}
