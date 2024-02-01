package main

import (
	"sync"
	"time"
)

/*
	为 sync.WaitGroup 中Wait函数支持 WaitTimeout 功能.
*/

type WaitTimeoutGroup struct {
	sync.WaitGroup
	timeout time.Duration
}

func NewWaitTimeoutGroup(timeout time.Duration) *WaitTimeoutGroup {
	return &WaitTimeoutGroup{
		timeout:   timeout,
		WaitGroup: sync.WaitGroup{},
	}
}

func (wtg *WaitTimeoutGroup) WaitTimeout() bool {
	waitC := make(chan bool, 1)
	// 使用 AfterFunc 计时，如果超过时间，传入 true
	go func() {
		time.AfterFunc(wtg.timeout, func() {
			waitC <- true
		})
	}()

	// 如果先返回就直接返回 false
	go func() {
		wtg.Wait()
		waitC <- false
	}()

	return <-waitC

}
