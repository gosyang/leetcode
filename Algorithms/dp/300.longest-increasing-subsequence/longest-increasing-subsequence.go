// Package _00_longest_increasing_subsequence - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* longest-increasing-subsequence - file brief introduce */
/*
modification history
-----------------------------
2020/2/3 1:09 PM, by gosyang, create
*/
/*
DESCRIPTION
给定一个无序的整数数组，找到其中最长上升子序列的长度。
说明:

可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
你算法的时间复杂度应该为 O(n2) 。
进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
*/
package _00_longest_increasing_subsequence

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 1
	res := 1
	for i := 1; i < n; i++ {
		// 只包含自己
		max := 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[j]+1 > max {
					max = dp[j]+1
				}
			}
		}
		dp[i] = max
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

// dp[i] 包含第i位的最长上升子序列
// dp[i] = max(dp[j]+1		, if nums[j] < nums[i]
//             1 			, else
//		  ) 0 <= j < i
// dp[0] = 1
// max(dp[i])
// TODO nlogn的二分查找的算法