package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
	限制 goroutine 数量
*/

func main() {
	// 限制 num 个 goroutine
	num := 5
	numC := make(chan struct{}, num)

	wg := sync.WaitGroup{}

	// 查看有多少个 goroutine
	go func() {
		for {
			num := runtime.NumGoroutine()
			fmt.Printf("Number of Goroutines: %d\n", num)
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		// 限制数量
		numC <- struct{}{}
		i := i
		go func(input int) {
			defer func() {
				wg.Done()
				// 完成前需要放回！
				<-numC
			}()
			fmt.Println(i)
			time.Sleep(time.Second)
		}(i)
	}
	wg.Wait()
}
