package main

/*
	2024. 考试的最大困扰度
	一位老师正在出一场由 n 道判断题构成的考试，每道题的答案为 true （用 'T' 表示）或者 false （用 'F' 表示）。老师想增加学生对自己做出答案的不确定性，方法是 最大化 有 连续相同 结果的题数。（也就是连续出现 true 或者连续出现 false）。

	给你一个字符串 answerKey ，其中 answerKey[i] 是第 i 个问题的正确结果。除此以外，还给你一个整数 k ，表示你能进行以下操作的最多次数：

	每次操作中，将问题的正确答案改为 'T' 或者 'F' （也就是将 answerKey[i] 改为 'T' 或者 'F' ）。
	请你返回在不超过 k 次操作的情况下，最大 连续 'T' 或者 'F' 的数目。



示例 1：

输入：answerKey = "TTFF", k = 2
输出：4
解释：我们可以将两个 'F' 都变为 'T' ，得到 answerKey = "TTTT" 。
总共有四个连续的 'T' 。
示例 2：

输入：answerKey = "TFFT", k = 1
输出：3
解释：我们可以将最前面的 'T' 换成 'F' ，得到 answerKey = "FFFT" 。
或者，我们可以将第二个 'T' 换成 'F' ，得到 answerKey = "TFFF" 。
两种情况下，都有三个连续的 'F' 。
示例 3：

输入：answerKey = "TTFTTFTT", k = 1
输出：5
解释：我们可以将第一个 'F' 换成 'T' ，得到 answerKey = "TTTTTFTT" 。
或者我们可以将第二个 'F' 换成 'T' ，得到 answerKey = "TTFTTTTT" 。
两种情况下，都有五个连续的 'T' 。
*/

// 采用滑动窗口实现此题，右指针使用 for 遍历，左指针按需移动

func maxConsecutiveAnswers(answerKey string, k int) int {
	// 分别传入 'T', 'F'
	return max(maxConsecutiveChar(answerKey, k, 'T'), maxConsecutiveChar(answerKey, k, 'F'))

}

func maxConsecutiveChar(answerKey string, k int, input byte) int {

	left := 0  // 左指针
	kCopy := k // 可变动条件
	res := 0   // 数量
	// 右指针使用遍历滑动
	for right := 0; right < len(answerKey); right++ {
		// 这里是当遍历的不是该元素时，需要扣除可变动条件
		if answerKey[right] != input {
			kCopy--
		}

		// 当馀额小于0时，需要移动左指针
		for kCopy < 0 {
			// 如果当下这个左指针元素是错的，记得要把馀额加回来
			if answerKey[left] != input {
				kCopy++
			}
			// 滑动左指针
			left++
		}
		// 比较大小一下。
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
