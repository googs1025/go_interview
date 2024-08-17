package main

import "fmt"

/*
	考初始化数组的区别 也要注意切片是共同引用底层数组，再改变时尽量使用 copy(x,x) 的方式
*/

func main() {
	s := make([]int, 5)
	ss := make([]int, 0, 5)

	s = append(s, 1, 2, 3)
	ss = append(ss, 1, 2, 3)

	fmt.Println(s)
	fmt.Println(ss)
}
