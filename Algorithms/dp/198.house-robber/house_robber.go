// Package _98_house_robber - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* house_robber - file brief introduce */
/*
modification history
-----------------------------
2020/1/29 12:50 PM, by gosyang, create
*/
/*
DESCRIPTION
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。

输入: [1,2,3,1]
输出: 4
解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。

*/
package _98_house_robber

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}
	max := -1
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = nums[1]
	if dp[0] > dp[1] {
		max = dp[0]
	} else {
		max = dp[1]
	}
	for i := 2; i < len(nums); i++ {
		m := -1
		for j := 0; j < i-1; j++ {
			if dp[j] > m {
				m = dp[j]
			}
		}
		dp[i] = m + nums[i]
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

// 上面这个这个方程适用于可以为负数的，类似最长子序列和
// a.dp[i]代表一定要偷第i家，多少钱
// b.dp[i] = max(dp[0]..dp[i-2])+nums[i]
// c. dp[0] = nums[0] dp[1] = nums[1]
// d. 求max(dp[i])

// 第二种方程，就是到第i家，最多多少钱，要么偷要么不偷, 假如有负数，那还应该比较 dp[i-2]如果是负的，直接nums[i]
// dp[i] = max(dp[i-2]+nums[i], dp[i-1])
func rob2(nums []int) int {
	preMax := 0
	curMax := 0
	for _, n := range nums {
		tmp := curMax
		if preMax+n > curMax {
			curMax = preMax + n
		}
		preMax = tmp
	}
	return curMax
}

// dfs 方法，全局记录最大值，然后遍历
var max int

func robDFS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max = -1
	dfs(nums, 0, false, 0)
	return max
}

// k 代表当前到第几位，choose 代表前一位有没有被选
func dfs(nums []int, k int, choose bool, sum int) {
	if k >= len(nums) {
		if sum > max {
			max = sum
		}
		return
	}
	if choose {
		// 前一位如果选了，k位不能选，跳到k+1
		dfs(nums, k+1, false, sum)
	} else {
		// 没选，k位可选，可不选
		dfs(nums, k+1, true, sum+nums[k])
		dfs(nums, k+1, false, sum)
	}
}
