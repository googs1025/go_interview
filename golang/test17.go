package main

import (
	"fmt"
	"strings"
)

/*
	判断字符串中字符是否全都不同
*/

func solution17(str string) bool {
	s := []rune(str)

	for _, v := range s {
		if strings.Count(str, string(v)) > 1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(solution17("adffssaa"))
}
