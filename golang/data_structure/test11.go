package main

import (
	"fmt"
	"sync"
)

/*
	实现一个 trylock 功能的对象
*/

type TryLock struct {
	ch chan struct{}
}

func NewTryLock() *TryLock {
	tryLock := &TryLock{ch: make(chan struct{}, 1)}
	// 重要，需要一开始初始化一次
	tryLock.ch <- struct{}{}
	return tryLock
}

func (t *TryLock) TryLock() bool {
	// 默认没抢到锁
	res := false
	select {
	// 当获取到 chan 值时，代表抢到锁
	case <-t.ch:
		res = true
	default:
	}
	return res
}

func (t *TryLock) UnLock() {
	// 回复
	t.ch <- struct{}{}
}

func main() {
	var wg sync.WaitGroup

	tryLock := NewTryLock()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if !tryLock.TryLock() {
				fmt.Println("get lock failed: ", i)
				return
			}
			fmt.Println("get lock: ", i)
			tryLock.UnLock()
		}(i)
	}

	wg.Wait()

}
