// Package main - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* unique_binary_search_trees - file brief introduce */
/*
modification history
-----------------------------
2020/1/31 5:17 PM, by gosyang, create
*/
/*
DESCRIPTION
给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？
*/
package main

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i < n+1; i++ {
		sum := 0
		for j := 0; j < i; j++ {
			sum += dp[j]*dp[i-1-j]
		}
		dp[i] = sum
	}
	return dp[n]
}

// 一看到题可以递归啊，考虑dp
// dp[i] 表示1...i个连续的数有几种二叉搜索树, x,x+1,....x+i-1
// dp[i] 可以是以x为根，到x+i-1为根，同时，首尾对称的是dp[i-1]个，因此
//       = (dp[0] * dp[i-1]) + (dp[1] * dp[i-2]) + ... + (dp[i-2] * dp[1]) + (dp[i-1] * dp[0])
// dp[0] = dp[1] = 1
// 求dp[n]

// 特别注意的是，假如i是从0开始的，那么j不会进循环，导致dp[0]会变成sum=0！！！
