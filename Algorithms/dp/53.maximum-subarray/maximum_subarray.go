// Package _3_maximum_subarray - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* maximum_subarray - file brief introduce */
/*
modification history
-----------------------------
2020/1/28 9:34 PM, by gosyang, create
*/
/*
DESCRIPTION
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），
返回其最大和。
*/
package _3_maximum_subarray

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1]+nums[i] > nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

// dp定义就是包含第i个数的最大子序列和，dp[i] = max(dp[i-1]+nums[i], nums[i])
// 结果就是max(dp[i]), 0<=i<len(nums)

// 优化空间的话你会发现dp没必要定义那么多，甚至直接覆盖nums也可以, 因此普通dp看有没有优化空间的方法

func maxSubArray2(nums []int) int {
	dp := nums[0]
	res := dp
	for i := 1; i < len(nums); i++ {
		if dp > 0 {
			dp += nums[i]
		} else {
			dp = nums[i]
		}
		if dp > res {
			res = dp
		}
	}
	return res
}

// 结合198题大街的考虑，dp[i]定义为前i个数的最大和，dp[i] = max(dp[i-1]+nums[i], dp[i-1], nums[i]),
// 其实其中的对dp[i-1]的比较，就是max(dp[i])的过程，这样，直接最后dp[len(nums)]即可
