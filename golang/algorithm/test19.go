package main

import "sort"

/*
	三数之和
	nums = [-1,0,1,2,-1,4] target = 0
    return [[-1,0,1],[-1,-1,2]]
*/

func threeSum(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)
	res := [][]int{}
	// 遍历
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			if (nums[i] + nums[left] + nums[right]) == 0 {
				aa := []int{}
				aa = append(aa, nums[i])
				aa = append(aa, nums[left])
				aa = append(aa, nums[right])
				res = append(res, aa)
				for (left < right) && nums[left] == nums[left+1] {
					left++
				}
				for (left < right) && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if (nums[i] + nums[left] + nums[right]) < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
