// Package _4_minimum_path_sum - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* minimum_path_sum - file brief introduce */
/*
modification history
-----------------------------
2020/1/29 9:53 PM, by gosyang, create
*/
/*
DESCRIPTION
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。
*/
package _4_minimum_path_sum

func minPathSum(grid [][]int) int {
	n := len(grid[0])

	dp := make([]int, n)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[0] = grid[0][0]
				continue
			}
			if i == 0 {
				dp[j] = dp[j-1] + grid[i][j]
				continue
			}
			if j == 0 || dp[j] < dp[j-1] {
				dp[j] += grid[i][j]
			} else {
				dp[j] = dp[j-1] + grid[i][j]
			}
		}
	}
	return dp[n-1]
}

// 直接上dp了
// a. dp[i,j] 代表从[0,0]走到此的最小和
// b. dp[i,j] = min(dp[i-1,j], dp[i,j-1]) + grid[i,j]
// c. 边界条件 dp[0,j] = grid[0,j], dp[i,0] = grid[i,0]
// d. 结果dp[m-1,n-1]
// e. 数据压缩下，不需要的空间去掉

// 和62题一样，最好用从0开始的思路来做，一样的思路
