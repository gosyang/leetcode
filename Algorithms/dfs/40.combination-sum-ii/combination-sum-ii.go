// Package _0_combination_sum_ii - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* combination-sum-ii - file brief introduce */
/*
modification history
-----------------------------
2020/4/19 10:39 PM, by gosyang, create
*/
/*
DESCRIPTION
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的每个数字在每个组合中只能使用一次。
说明：
所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。 
示例 1:
输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:
输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]
*/
package _0_combination_sum_ii

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	DFS(candidates, target, []int{}, 0, &res)
	return res
}

func DFS(candidates []int, target int, cur []int, left int, res *[][]int) {
	if target == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for i := left; i < len(candidates); i++ {
		if target - candidates[i] < 0 {
			continue
		}
		// 这里很精髓，意味着如果是利用重复元素填充后几位没问题，但是利用重复的元素填第一位就不行
		if i > left && candidates[i] == candidates[i-1] {
			continue
		}
		// i+1 保证每个元素只用一次
		DFS(candidates, target-candidates[i], append(cur, candidates[i]), i+1, res)
	}
}