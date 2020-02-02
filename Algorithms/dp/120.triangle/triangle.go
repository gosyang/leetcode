// Package _20_triangle - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* triangle - file brief introduce */
/*
modification history
-----------------------------
2020/2/2 7:54 PM, by gosyang, create
*/
/*
DESCRIPTION
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

说明：
如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
*/
package _20_triangle

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n+1)

	for i := n-1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if dp[j] < dp[j+1] {
				dp[j] += triangle[i][j]
			}  else {
				dp[j] = dp[j+1] + triangle[i][j]
			}
		}
	}
	return dp[0]
}


// a. dp[i][j] 代表走到i,j的路径最小和, 自底向上比较好推，
// b. 转移方程： dp[i][j] = min(dp[i+1][j], dp[j+1][i+1]) + triangle[i][j]
// c. 边界， dp[last][j] = triangle[last][j]
// d. 求， dp[0][0]
// e. 压缩状态，dp[i]为上一行的内容即可
// 注意点： 1. dp要多加1，因为要考虑最右下角边界，2. j的上界是i