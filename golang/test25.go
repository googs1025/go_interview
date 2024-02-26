package main

import "fmt"

type People struct{}

func (p *People) MethodA() {
	fmt.Println("methodA")
	p.MethodB()
}

func (p *People) MethodB() {
	fmt.Println("methodB")
}

type Teacher struct {
	// 组合
	People
	//p People  这就不是组合
}

func (t *Teacher) MethodB() {
	fmt.Println("Teacher methodB")
}

func main() {
	t := &Teacher{}
	t.MethodB()
	// 虽然 t 没有直接实现 MethodA 方法，但有struct组合，会去往上找有实现的对象去执行方法
	t.MethodA()
}
