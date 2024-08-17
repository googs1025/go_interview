package main

import "fmt"

/*
	注意 接口类型实现的方法集是用指针类型，那申明接口时也需要使用指针类型
*/

type People1 interface {
	Speak(str string) string
}

type student1 struct {
}

func (s *student1) Speak(str string) string {
	if str == "test" {
		return "test"
	}
	return "no"
}

func main() {
	var peo People1 = &student1{}
	//var peo People1 = student1{}
	fmt.Println(peo.Speak("test"))
}
