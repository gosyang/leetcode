// Package _6_permutations - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* permutations - file brief introduce */
/*
modification history
-----------------------------
2020/4/19 10:49 PM, by gosyang, create
*/
/*
DESCRIPTION
给定一个 没有重复 数字的序列，返回其所有可能的全排列。
*/
package _6_permutations

func permute(nums []int) [][]int {
	res := [][]int{}
	DFS(nums, 0, &res)
	return res
}

func DFS(nums []int, n int, res *[][]int) {
	if n == len(nums) {
		tmp := make([]int, n)
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}
	for i := n; i < len(nums); i++ {
		nums[i], nums[n] = nums[n], nums[i]
		DFS(nums, n+1, res)
		nums[i], nums[n] = nums[n], nums[i]
	}
}