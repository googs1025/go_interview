package main

import "fmt"

/*
	实现一个set
*/

type Set struct {
	data map[interface{}]struct{}
}

func NewSet() *Set {
	return &Set{
		data: map[interface{}]struct{}{},
	}
}

func (s *Set) Add(key interface{}) {
	s.data[key] = struct{}{}
}

func (s *Set) Get(key interface{}) (interface{}, bool) {
	if v, ok := s.data[key]; ok {
		return v, true
	}

	return nil, false
}

func (s *Set) Delete(key interface{}) {
	delete(s.data, key)
}

func main() {
	s := NewSet()
	s.Add("aaa")
	s.Add("aaa")
	fmt.Println(s.Get("aaa"))
}
