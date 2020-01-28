// Package _0_climbing_stairs - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* climbing_stairs - file brief introduce */
/*
modification history
-----------------------------
2020/1/28 9:11 PM, by gosyang, create
*/
/*
DESCRIPTION
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。
*/
package _0_climbing_stairs

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 比较两种dp， 第一个是对边界的处理，，第二种是逻辑更清晰一点，不考虑第0层啥意思
func climbStairs2(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
