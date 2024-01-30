package main

import (
	"fmt"
	"time"
)

/*
	函数设置超时时间控制
*/

func task() chan int {
	res := make(chan int)
	go func() {
		time.Sleep(time.Second * 2)
		res <- 100
	}()
	return res
}

func run() int {

	resC := task()

	select {
	case res := <-resC:
		return res
	case <-time.After(time.Second * 1):
		fmt.Println("timeout")
		return 0
	}

}

func main() {
	run()
}
