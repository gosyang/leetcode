/* Copyright
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/
package main

import "strings"

// 滑动窗口，用每次最右边的和已有的比，有重复的，就把左边界移动 +index+1
func lengthOfLongestSubstring(s string) int {
	ans := 0
	left := 0
	right := 0
	s1 := s[left:right]
	for ; right < len(s); right++ {
		if index := strings.IndexByte(s1, s[right]); index != -1 {
			left += index + 1
		}
		s1 = s[left : right+1]
		if len(s1) > ans {
			ans = len(s1)
		}
	}
	return ans
}
