// Package _3_unique_paths_ii - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* unique_paths_ii - file brief introduce */
/*
modification history
-----------------------------
2020/1/30 10:57 AM, by gosyang, create
*/
/*
DESCRIPTION
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

网格中的障碍物和空位置分别用 1 和 0 来表示。
说明：m 和 n 的值均不超过 100。
*/
package _3_unique_paths_ii

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid[0])
	dp := make([]int, n)
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				continue
			}
			if i == 0 && j == 0 {
				dp[j] = 1
				continue
			}
			if i == 0 {
				dp[j] = dp[j-1]
			} else if j == 0 {
				dp[j] = dp[j]
			} else {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n-1]
}

// a. dp[i,j] 到[i,j]的不同路径总数
// b. dp[i,j] = dp[i-1,j] + dp[i,j-1], 如果obstacleGrid[i,j]=1 则dp[i,j]==1
// c. dp[0,j] = 1, dp[i,0] = 1 ！！ 这里不对，如果第一行第一列里有1 的，这一行和列都不行了
// d. 求dp[m-1][n-1]
// e. 数据压缩

// 这题掉坑里了，边界条件不对
