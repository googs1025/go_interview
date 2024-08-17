package main

import (
	"fmt"
	"time"
)

/*
	写出以下逻辑，要求每秒钟调用一次proc并保证程序不退出?

	注意：panic 使用 defer+recover() 不能在父goroutine来恢复，只能在该goroutine恢复而已，需要特别注意
*/

func main() {

	go func() {
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc函数
		// 3 要求程序不能退出

		for {
			select {
			case <-time.After(time.Second):

				go func() {
					defer func() {
						if r := recover(); r != nil {
							fmt.Println("aaa")
						}
					}()
					proc()
					fmt.Println("调用～～")
				}()
			}

		}

	}()

	fmt.Println("阻塞在这里～")
	select {}
}

func proc() {
	panic("ok")
}
