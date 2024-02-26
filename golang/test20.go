package main

import "fmt"

/*
	主要考察 defer 与 panic 知识点
	panic 会等 defer结束后才向上传递。
	output:
	打印后
	打印中
	打印前
	panic: 触发异常
*/

func main() {
	// 第一题
	//deferCall()

	// 第二题
	//a, b := 1, 0
	//defer calc("1", a, calc("10", a, b))
	//defer calc("2", a, calc("20", a, b))

	// 第三题
	fmt.Println(DeferFunc1(1))
	fmt.Println(DeferFunc2(1))
	fmt.Println(DeferFunc3(1))
}

// 第一题
func deferCall() {
	defer func() {
		fmt.Println("打印前")
	}()
	defer func() {
		fmt.Println("打印中")
	}()

	defer func() {
		fmt.Println("打印后")
	}()
	panic("触发异常")
}

// 第二题
func calc(index string, a, b int) int {
	res := a + b
	fmt.Println("index: ", index, "res: ", res)
	return res
}

func DeferFunc1(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) int {
	t := i
	t += i
	return t
}
