package main

import "fmt"

/*
	select 的随机性，有时候会panic 有时候不会。
*/

func main() {
	intCh := make(chan int, 1)
	strCh := make(chan string, 1)
	intCh <- 1
	strCh <- "test"

	select {
	case v := <-strCh:
		fmt.Println(v)
	case <-intCh:
		panic("chan panic")
	}
}
