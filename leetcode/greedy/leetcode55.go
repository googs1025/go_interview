package main

import "fmt"

/*
	给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。

	每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:

	0 <= j <= nums[i]
	i + j < n
	返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。



	示例 1:

	输入: nums = [2,3,1,1,4]
	输出: 2
	解释: 跳到最后一个位置的最小跳跃数是 2。
		 从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
	示例 2:

	输入: nums = [2,3,0,1,4]
	输出: 2
*/

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{2, 3, 0, 1, 4}))
}

func jump(nums []int) int {
	// 记录多少步
	res := 0
	// 现在这步能跳的最远距离
	curDistance := 0
	// 下一步能跳的最远距离
	nextDistance := 0

	for i := 0; i < len(nums); i++ {
		// 比较我目前这步的最远距离和我下一步的最远距离谁远
		nextDistance = max2(nextDistance, nums[i]+i)
		// 当我这步已经是最遠，我只能
		if i == curDistance {
			if curDistance != len(nums)-1 {
				res++
				curDistance = nextDistance
				if nextDistance > len(nums)-1 {
					break
				}
			}
		}
	}
	return res

}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
