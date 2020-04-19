// Package _7_permutations_ii - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* permutations-ii - file brief introduce */
/*
modification history
-----------------------------
2020/4/19 10:47 PM, by gosyang, create
*/
/*
DESCRIPTION
xx给定一个可包含重复数字的序列，返回所有不重复的全排列。
*/
package _7_permutations_ii

func permuteUnique(nums []int) [][]int {
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
	// 对于包含重复数字的排列，存在例如 1,1,2 两个1 交换但其实一样的情况，这种就需要在第二个1 做头的时候剪掉
	m := map[int]bool{}
	for i := n; i < len(nums); i++ {
		if m[nums[i]] {
			continue
		}
		m[nums[i]] = true
		nums[i], nums[n] = nums[n], nums[i]
		DFS(nums, n+1, res)
		nums[i], nums[n] = nums[n], nums[i]
	}
}