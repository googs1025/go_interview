package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	优胜劣汰: 下发多个 goroutine 但只取最快返回的结果
*/

func job() int {
	res := rand.Intn(20)
	time.Sleep(time.Second * time.Duration(res))
	return res
}

func main() {

	num := 10
	chanC := make(chan int, num)

	for i := 0; i < 5; i++ {
		go func() {
			chanC <- job()
		}()
	}

	fmt.Println("使用最快的 goroutine: ", <-chanC, "秒")
}
