package main

import (
	"fmt"
	"time"
)

/*
	协程实现顺序打印123
*/

func main() {

	oneC := make(chan struct{})
	twoC := make(chan struct{})
	threeC := make(chan struct{})

	go func() {
		time.Sleep(time.Second)
		oneC <- struct{}{}
	}()

	for i := 0; i < 5; i++ {
		go func() {
			<-oneC
			fmt.Println("1")
			twoC <- struct{}{}
		}()

		go func() {
			<-twoC
			fmt.Println("2")
			threeC <- struct{}{}
		}()

		go func() {
			<-threeC
			fmt.Println("3")
			oneC <- struct{}{}
		}()
	}

	time.Sleep(time.Second * 5)
}



