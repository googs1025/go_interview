package main

import (
	"fmt"
	"time"
)

/*
	goroutine 交替打印
	goroutine a 打印 "A"
	goroutine b 打印 "B"
	goroutine c 打印 "C"

	输出： "ABCABCABCABCABC"

*/

func main() {
	res := ""

	chanA := make(chan struct{})
	chanB := make(chan struct{})
	chanC := make(chan struct{})
	go func() {
		time.Sleep(time.Second)
		// 需要触发第一个 goroutine
		chanA <- struct{}{}
	}()

	for i := 0; i < 5; i++ {
		// 等待第一个 goroutine 执行完毕
		<-chanA
		go func() {
			res += "A"
			// 需要触发第二个 goroutine
			chanB <- struct{}{}
		}()
		// 等待第一个 goroutine 执行完毕
		<-chanB
		go func() {
			res += "B"
			// 需要触发第三个 goroutine
			chanC <- struct{}{}
		}()
		// 等待第二个 goroutine 执行完毕
		<-chanC
		go func() {
			res += "C"
			// 需要触发第一个 goroutine
			chanA <- struct{}{}
		}()

	}

	time.Sleep(time.Second)
	fmt.Println(res)

}
