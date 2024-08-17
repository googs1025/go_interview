package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
	可以观察是否发生 goroutine泄漏
*/

type query func(string) string

func exec(name string, vs ...query) string {
	//ch := make(chan string)
	ch := make(chan string, 4)
	fn := func(i int) {
		ch <- vs[i](name)
	}
	for i, _ := range vs {
		go fn(i)
	}
	return <-ch
}

func main() {
	res := exec("111", func(s string) string {
		return s + "func1"
	}, func(s string) string {
		return s + "func2"
	}, func(s string) string {
		return s + "func3"
	}, func(s string) string {
		return s + "func4"
	})
	fmt.Println(res)

	for {
		select {
		case <-time.After(time.Second):
			fmt.Println("goroutine num: ", runtime.NumGoroutine())
		}
	}

}
