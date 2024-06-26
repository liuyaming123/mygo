package goRoutinePool

import (
	"fmt"
	"math/rand"
	"time"
)

type Job struct {
	Id      int
	RandNum int
}

type Result struct {
	job *Job
	// 求和
	sum int
}

func MyJob() {
	// 需要两个管道
	// 1.job管道
	jobChan := make(chan *Job, 128)
	// 2.结果管道
	resultChan := make(chan *Result, 128)
	// 3.创建工作池
	createPool(64, jobChan, resultChan)

	// 4.开个打印的协程
	go func(resultChan chan *Result) {
		for result := range resultChan {
			// 遍历结果管道打印
			fmt.Printf("job id: %v, randnum: %v, result: %d\n", result.job.Id,
				result.job.RandNum, result.sum)
		}
	}(resultChan)

	var id int
	for id < 5000000 {
		id++
		// 生成随机数
		r_num := rand.Int()

		job := &Job{
			Id:      id,
			RandNum: r_num,
		}

		jobChan <- job
	}

	time.Sleep(time.Second)
}

// 创建工作池
// 参数1：开几个协程
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	// 根据开协程个数去跑运行
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			// 遍历job所有数据进行相加
			for job := range jobChan {
				// 随机数接过来
				r_sum := job.RandNum
				// 随机数每一位相加
				// 定义返回值
				var sum int
				for r_sum != 0 {
					tmp := r_sum % 10
					sum += tmp
					r_sum /= 10
				}
				// 想要的结果是Result
				r := &Result{
					job: job,
					sum: sum,
				}
				// 运算结果扔到管道
				resultChan <- r
			}
		}(jobChan, resultChan)
	}

}
