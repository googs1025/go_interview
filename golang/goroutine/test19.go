package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 主要思考是否有初始化 chan ，如果没有初始化，默认chan都会阻塞
	var ch chan int
	go func() {
		// 打开此开关尝试，并且可以尝试 无缓冲和有缓冲的区别
		//ch = make(chan int)
		//ch = make(chan int, 1)
		ch <- 1
		fmt.Println("我结束了")
	}()

	go func() {
		time.Sleep(time.Second)
		<-ch
		fmt.Println("我结束了！！！")
	}()

	c := time.Tick(time.Second * 2)
	for range c {
		fmt.Printf("#goroutine: %d\n", runtime.NumGoroutine())
	}
}
