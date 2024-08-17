package main

import (
	"fmt"
	"strings"
)

/*
	判断两个给定的字符串排序后是否一致

	给定两个 str , 判断排序后会相同

	(大小写是不同字符，如果视为相同，那就先转为小写就行)

*/

func solution13(str1 string, str2 string) bool {
	s1 := []rune(str1)
	s2 := []rune(str2)

	if len(s1) != len(s2) {
		return false
	}

	for _, v := range s1 {
		if strings.Count(str1, string(v)) != strings.Count(str2, string(v)) {
			return false
		}
	}

	return true

}

func main() {
	fmt.Println(solution13("Abc", "cba"))
}
