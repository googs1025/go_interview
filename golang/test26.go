package main

import "sync"

// 涉及到 并发场景的map 需要使用 lock 或是 sync.Map 对象实现

func main() {

}

type UserAge struct {
	ages map[string]int
	sync.Mutex
}

func (u *UserAge) Add(name string, age int) {
	u.Lock()
	defer u.Unlock()
	u.ages[name] = age
}

func (u *UserAge) Get(name string) int {
	u.Lock()
	defer u.Unlock()
	if age, ok := u.ages[name]; ok {
		return age
	}
	return -1
}
