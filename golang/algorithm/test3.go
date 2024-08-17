package main

import "fmt"

/*
	只出现一次的数字
	ex: nums = [1, 2, 3, 3, 2, 1, 5]
	return = [5]
*/

func main() {

	nums := []int{1, 2, 3, 3, 2, 1, 5}
	fmt.Println(Solution3_1(nums))
}

/*
	有两种方法：
	1. map 方式遍例记录，再遍例一次取出为1的结果
	2. 异或运算
*/

func Solution3_1(nums []int) int {

	hash := map[int]int{}

	for _, v := range nums {
		hash[v]++
	}

	for i := range hash {
		if hash[i] == 1 {
			return i
		}
	}
	return 0
}

func Solution3(nums []int) int {

	var res int
	for _, v := range nums {
		res = res ^ v
	}
	return res
}
