// Package _39_word_break - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* word_break - file brief introduce */
/*
modification history
-----------------------------
2020/2/2 8:47 PM, by gosyang, create
*/
/*
DESCRIPTION
给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：
拆分时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。
*/
package _39_word_break

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		tmp := false
		for _, w := range wordDict {
			if i-len(w) < 0 {
				continue
			}
			if s[i-len(w):i] == w {
				tmp = tmp || dp[i-len(w)]
			}
		}
		dp[i] = tmp
	}
	return dp[len(s)]
}

// 啥搜肯定不行，有最优子结构，即最后几个字符如果有某个单词，那其实就只判断前面就可以了
// dp[i] 表示s[:i+1] 是否可以被拆分
// dp[i] = dp[:i-len(word1)] || dp[:i-len(word2)]....有匹配结尾的 or 全不匹配的话 false
// dp[0] = true
// 求 dp[len(s)]
