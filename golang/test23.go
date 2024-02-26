package main

import (
	"fmt"
	"sync"
)

/*
	注意闭包和传入时 goroutine的变化
*/

func main() {
	var wg sync.WaitGroup

	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}

	for i := 10; i < 20; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println("i: ", i)
		}(i)
	}

	wg.Wait()

}
