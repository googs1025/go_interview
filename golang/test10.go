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
		chanA <- struct{}{}
	}()

	for i := 0; i < 5; i++ {
		<-chanA
		go func() {
			res += "A"
			chanB <- struct{}{}
		}()
		<-chanB
		go func() {
			res += "B"
			chanC <- struct{}{}
		}()
		<-chanC
		go func() {
			res += "C"
			chanA <- struct{}{}
		}()

	}

	time.Sleep(time.Second)
	fmt.Println(res)

}
