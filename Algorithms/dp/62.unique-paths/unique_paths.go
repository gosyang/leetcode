// Package _2_unique_paths - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* unique_paths - file brief introduce */
/*
modification history
-----------------------------
2020/1/29 10:42 PM, by gosyang, create
*/
/*
DESCRIPTION
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
问总共有多少条不同的路径？
说明：m 和 n 的值均不超过 100。
*/
package _2_unique_paths

func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[j] = 1
			} else {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n-1]
}

// dp[i,j] 表示到[i,j]的路径数
// dp[i,j] = dp[i-1,j] + dp[i, j-1]
// dp[0,j] = dp[i,0] = 1, 原理要吃透是因为旁边和上面没了，是有一边有，不能武断是1，真正的边界只有dp[0,0]=1
// 最终为dp[m-1,n-1]
// 压缩下成一维，并对于0的情况直接在循环赋值即可
// TODO 【注意】这道题这种解法是在没有障碍的情况下，可以直接全是1, 最好用通用的办法，第一行第一列也是公式赋值的好

func uniquePaths2(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[j] = 1
				continue
			}
			if i == 0 {
				dp[j] = dp[j-1]
			} else if j == 0 {
				continue
			} else {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n-1]
}