package main

/*
	1004. 最大连续1的个数 III
	给定一个二进制数组 nums 和一个整数 k，如果可以翻转最多 k 个 0 ，则返回 数组中连续 1 的最大个数 。



示例 1：

输入：nums = [1,1,1,0,0,0,1,1,1,1,0], K = 2
输出：6
解释：[1,1,1,0,0,1,1,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 6。
示例 2：

输入：nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
输出：10
解释：[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 10。
*/

func longestOnes(nums []int, k int) int {
	// 答案
	res := 0
	// 左指针
	left := 0
	// 右指针直接滑动
	for right := 0; right < len(nums); right++ {
		// 如果不是 1, 代表需要扣除馀额
		if nums[right] != 1 {
			k--
		}
		//
		for k < 0 {
			// 如果恰好这个数是 0, 代表需要增加馀额
			if nums[left] != 1 {
				k++
			}
			// 滑动左指针
			left++
		}
		// 替换
		res = max(res, right-left+1)
	}
	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
