package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	写出 生产者消费者模式
*/

func producer(inputC chan int) {

	defer close(inputC)

	for i := 0; i < 10; i++ {

		inputC <- rand.Int()
		time.Sleep(time.Second)
	}
}

func consumer(inputC chan int, stopC chan struct{}) {

	for i := range inputC {
		i := i
		go func(num int) {
			fmt.Println(num)
		}(i)
	}

	stopC <- struct{}{}

}

func main() {
	chanC := make(chan int)
	stopC := make(chan struct{})
	go producer(chanC)
	go consumer(chanC, stopC)
	<-stopC
}
