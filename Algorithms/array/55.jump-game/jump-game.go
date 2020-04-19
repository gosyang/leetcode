// Package _5_jump_game - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* jump-game - file brief introduce */
/*
modification history
-----------------------------
2020/4/17 12:43 PM, by gosyang, create
*/
/*
DESCRIPTION
给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个位置。

示例 1:
输入: [2,3,1,1,4]
输出: true
解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。

示例 2:
输入: [3,2,1,0,4]
输出: false
解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
*/
package _5_jump_game

func canJump(nums []int) bool {
	reach := map[int]bool{}
	last := len(nums)-1
	reach[0] = true
	for i := 0; i < last; i++ {
		if !reach[i] {
			return false
		}
		// 如果本次怎么跳都不如之前，那就别循环了
		if i > 0 && nums[i] < nums[i-1] {
			continue
		}
		for j := 1; j <= nums[i]; j++ {
			r := i + j
			if r == last {
				return true
			}
			reach[r] = true
		}
	}
	return reach[last]
}

// 改进的话，其实如果i+nums[i]就是最远距离，