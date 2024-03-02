package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	两个协程交替打印一个数组，使数组中的数据按顺序输出
*/

var wg sync.WaitGroup

func main() {
	ch := make(chan struct{})
	inputC := make(chan string)
	stopC := make(chan struct{})
	s1 := []string{"a", "b", "c", "d", "e", "f"}

	go func() {
		for i := 0; i < len(s1); i++ {
			inputC <- s1[i]
		}
	}()

	go work1(inputC, ch, stopC)
	go work2(inputC, ch, stopC)

	time.Sleep(time.Second * 20)

}

func work1(input chan string, ch chan struct{}, stopC chan struct{}) {

	for {
		select {
		case <-ch:
			v := <-input
			fmt.Println(v)
			time.Sleep(time.Second)
			ch <- struct{}{}
		case <-stopC:
			return
		}
	}
}

func work2(input chan string, ch chan struct{}, stopC chan struct{}) {
	ch <- struct{}{}

	for {
		select {
		case <-ch:
			v := <-input
			fmt.Println(v)
			time.Sleep(time.Second)
			ch <- struct{}{}
		case <-stopC:
			return
		}
	}
}
