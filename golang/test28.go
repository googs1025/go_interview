package main

import "fmt"

/*
	接口对象：值+类型 所以不会 nil==nil
*/

type People2 interface {
	Show()
}

type student2 struct{}

func (s *student2) Show() {

}

func live() People2 {
	var stu *student2
	return stu
}

func main() {
	if live() == nil {
		fmt.Println("nil==nil")
	} else {
		fmt.Println("nil!=nil")
	}
}
