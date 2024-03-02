package main

import (
	"fmt"
	"time"
)

/*
	一个线程打印奇数一个线程打印偶数 交替打印
*/

func main() {
	ch := make(chan int)
	go goroutine1(ch)
	go goroutine2(ch)
	time.Sleep(time.Second)
}

func goroutine1(ch chan int) {
	for i := 0; i <= 20; i++ {
		ch <- i
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func goroutine2(ch chan int) {
	for i := 1; i <= 20; i++ {
		v := <-ch
		if v%2 == 1 {
			fmt.Println(v)
		}
	}
}
