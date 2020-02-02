// Package __longest_palindromic_substring - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* longest_palindromic_substring - file brief introduce */
/*
modification history
-----------------------------
2020/1/31 12:44 PM, by gosyang, create
*/
/*
DESCRIPTION
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。


*/
package __longest_palindromic_substring

func longestPalindrome(s string) string {
	if s == "" || len(s) == 1 {
		return s
	}
	var dp [1000][1000]bool
	for i := 0; i < 1000; i++ {
		dp[i][i] = true
	}
	// 左边界 和 右边界
	res := s[:1]
	max := 1
	// 右是j， 左是i
	for j := 1; j < len(s); j++ {
		for i := j - 1; i >= 0; i-- {
			if j-i == 1 {
				dp[i][j] = s[i] == s[j]
			} else if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = false
			}
			// 记录
			if dp[i][j] {
				if j-i+1 > max {
					max = j - i + 1
					res = s[i : j+1]
				}
			}
		}
	}
	return res
}

// dp
// a. 定义状态  dp[i][j] s的i到j位是否是回文子串, bool类型
// b. 转移方程  dp[i][j] = 1. 如果s[i]==s[j], dp[i][j] = dp[i+1][j-1]
//                        2. 如果s[i]!=s[j], dp[i][j] = 0
// c. 边界 i==j dp[i][j] = 1
// d. 结果 记录的最长
// e. 数据压缩先不压缩的话比较差
// f. 易错点，res的初始值要为s[:1]即第一个字符, 而不是空的, 不能res是"",max=1
// TODO O(1)的空间复杂度, 什么中心扩展法
