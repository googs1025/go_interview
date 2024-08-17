package main

import (
	"errors"
	"fmt"
)

func main() {
	// <nil>
	// f2 error
	// <nil>
	f1()
	f2()
	f3()
}

func f1() {
	var err error
	defer fmt.Println(err)
	err = errors.New("f1 error")
	return
}

func f2() {
	var err error
	defer func() {
		fmt.Println(err)
	}()
	err = errors.New("f2 error")
	return
}

func f3() {
	var err error
	defer func(err error) {
		fmt.Println(err)
	}(err)
	err = errors.New("f3 error")
	return
}
