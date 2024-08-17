package main

import "fmt"

/*
	两数之和：
	ex: 给定数组：nums = [2,7,11,15]  -> target = 9
	return [0, 1], 因为下标 0, 1 数字和为 9
*/

func main() {
	fmt.Println(twoSum1())
}

func twoSum() []int {
	var res []int

	num := []int{2, 7, 11, 15}

	target := 9

	for k, v := range num {
		for i := k + 1; i < len(num); i++ {
			if target-v == num[i] {
				res = append(res, k)
				res = append(res, i)
			}
		}
	}

	return res
}

func twoSum1() []int {
	var res []int

	num := []int{2, 7, 11, 15}
	hash := make(map[int]int)
	target := 9

	for k, v := range num {
		if value, ok := hash[target-v]; ok {
			res = append(res, value)
			res = append(res, k)
		}
		hash[v] = k
	}

	return res
}
