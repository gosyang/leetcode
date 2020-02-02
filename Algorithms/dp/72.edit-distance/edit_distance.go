// Package _2_edit_distance - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* edit_distance - file brief introduce */
/*
modification history
-----------------------------
2020/1/31 3:13 PM, by gosyang, create
*/
/*
DESCRIPTION
给定两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数 。
你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符

输入: word1 = "horse", word2 = "ros"
输出: 3
解释:
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
*/
package _2_edit_distance

func minDistance(word1 string, word2 string) int {
	l1 := len(word1) + 1
	l2 := len(word2) + 1

	dp := make([]int, l2)

	// 第0行，如果单词1是空，那么单词2是几个就几个操作
	for i := 0; i < l2; i++ {
		dp[i] = i
	}

	for i := 1; i < l1; i++ {
		// 需要比较0的下标
		old := dp[0]
		// 第 i 行的第一个，代表到新的一行，单词1i个，单词2 0个，需要i步
		dp[0] = i
		for j := 1; j < l2; j++ {
			tmp := dp[j]
			if word1[i-1] == word2[j-1] {
				dp[j] = old
			} else {
				dp[j] = min(min(dp[j], dp[j-1]), old) + 1
			}
			old = tmp
		}
	}
	return dp[l2-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// a.状态定义 dp[i][j]代表单词1到i位，2到j位相同，最少操作数
// b.转移方程 dp[i][j] = dp[i-1][j-1] 如果字母相同 or
//                      min(dp[i-1][j-1], 如果不同，替换最后一位
//                      dp[i-1][j], 删除1位
//                      dp[i][j-1], 增加1位) + 1
// c.边界条件 dp[0][j] = j, dp[i][0] = i
// d. 结果 dp[len(word1)][len(word2)]
// e. 数据压缩，只用到i-1行, j-1的新旧值都得记住，比较繁琐 TODO 下次再默写一遍
// 结合91题数字-字母解码的tmp回味
