package gotest

import (
	"fmt"
	"sync"
	"testing"
)

func TestM1(t *testing.T) {
	done := make(chan int, 10)
	go func() {
		fmt.Println("haha")
		done <- 5 // 发送
	}()
	x := <-done // 接收,不写接收者x表示丢弃接收
	fmt.Println(x)
	fmt.Println(cap(done))
}

func TestM2(t *testing.T) {
	done := make(chan int, 10)

	for i := 0; i < cap(done); i++ {
		fmt.Println("hehe")
		done <- 1
	}

	// 等待N个后台线程完成
	for i := 0; i < cap(done); i++ {
		<-done
	}
}

func TestM3(t *testing.T) {
	var wg sync.WaitGroup

	// 开N个后台打印线程
	// wg.Add(4)
	for i := 0; i < 100; i++ {
		wg.Add(1) // 增加等待事件的个数，必须确保在后台线程启动之前执行

		go func() {
			fmt.Println("hello, world!")
			wg.Done() // 表示完成一个事件
		}()
	}

	wg.Wait() // 等待全部的事件完成
}

// var wg sync.WaitGroup
// wg.Add(maxSnapNum)
// for i := 0; i < maxSnapNum; i++ {
// 	go performSnap(i, shotRate, outInfo.SnapshotInfo.StartTime, snapShotCmd, filePath, outputPath, &wg, tc.task.Vid)
// }
// wg.Wait()
