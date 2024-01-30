package main

import (
	"fmt"
	"strings"
)

/*
	判断回文数: 从前面遍例和后面遍例一致
*/

func main() {
	words := []string{"1234", "12221", "dad", "dsa cds", "ds  sd"}
	for _, v := range words {
		fmt.Println(Solution(v))
	}
}

func Solution(str string) bool {

	str = strings.ReplaceAll(str, " ", "")
	strList := strings.Split(str, "")
	var i, j = 0, len(str) - 1

	for i < j {
		if strList[i] != strList[j] {
			return false
		}
		i++
		j--
	}
	return true
}
