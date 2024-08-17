package main

import (
	"fmt"
	"time"
)

/*
	一个线程打印奇数一个线程打印偶数 交替打印
*/

func main() {
	//ch := make(chan struct{})
	////go goroutine1(ch)
	////go goroutine2(ch)
	//go aa(ch, "A")
	//go bb(ch, "B")
	//time.Sleep(time.Second)

	aaaa()
	select {}
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

func aaaa() {

	chanA := make(chan struct{})
	chanB := make(chan struct{})
	chanC := make(chan struct{})

	go func() {
		fmt.Println("start...")
		chanA <- struct{}{}
	}()

	for {
		go func() {
			<-chanA
			fmt.Println("A")
			time.Sleep(time.Second)
			chanB <- struct{}{}

		}()

		go func() {
			<-chanB
			fmt.Println("B")
			time.Sleep(time.Second)
			chanC <- struct{}{}
		}()

		go func() {
			<-chanC
			fmt.Println("C")
			time.Sleep(time.Second)
			chanA <- struct{}{}
		}()
	}
}
