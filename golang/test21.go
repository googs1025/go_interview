package main

import "fmt"

/*
	主要考 for range 时，都是值传递，
	请使用 指定数组下标的方式
*/

type Student struct {
	name string
	age  int
}

func student() {
	m := make(map[string]*Student)
	stus := []Student{
		{name: "test1", age: 20},
		{name: "test2", age: 20},
	}

	for k, _ := range stus {
		//m[sts.name] = &sts
		m[stus[k].name] = &stus[k]
	}

	for _, v := range m {
		fmt.Println(v.name)
	}
}

func main() {
	student()
}
