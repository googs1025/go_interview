package main

import (
	"sync"
	"time"
)

/*
	实现一个带过期时间且并发执行的map
*/

type TimeoutMap struct {
	data    sync.Map
	timeout time.Duration
}

func NewTimeoutMap(timeout time.Duration) *TimeoutMap {
	return &TimeoutMap{
		data:    sync.Map{},
		timeout: timeout,
	}
}

func (tm *TimeoutMap) Add(key interface{}, value interface{}) {
	tm.data.Store(key, value)

	defer func() {
		time.AfterFunc(tm.timeout, func() {
			tm.data.Delete(key)
		})
	}()
}

func (tm *TimeoutMap) Get(key interface{}) (interface{}, bool) {
	return tm.data.Load(key)
}

func main() {
	tm := NewTimeoutMap(time.Duration(5))
	tm.Add("aa", "aaa")
	tm.Add("ddd", "ddd")

	tm.Get("aaa")
}

