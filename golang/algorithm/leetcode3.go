package main

import "fmt"

/*
    给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
	示例 1:

	输入: s = "abcabcbb"
	输出: 3
	解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

	示例 2:

	输入: s = "bbbbb"
	输出: 1
	解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

	示例 3:

	输入: s = "pwwkew"
	输出: 3
	解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
		 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(str string) int {
	ans := make([]byte, 0)
	length := 0

	for i := 0; i < len(str); i++ {
		// 如果 str[i] 不在 ans 中, 直接加入, 并更新 length 长度
		if index, ok := isInAnsAndIndex(str[i], ans); !ok {
			ans = append(ans, str[i])
			length = max(length, len(ans))
		} else {
			// 如果 str[i] 在 ans 中，需要把 ans 之前的删除，
			// 并把 str[i] 加入
			ans = ans[index+1:]
			ans = append(ans, str[i])
		}
	}
	return length

}

// isInAns 查找是否 input 在数组中，并返回 index
func isInAnsAndIndex(input byte, ans []byte) (int, bool) {
	if len(ans) == 0 {
		return 0, false
	}
	//
	for i := 0; i < len(ans); i++ {
		if input == ans[i] {
			return i, true
		}
	}
	return 0, false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
