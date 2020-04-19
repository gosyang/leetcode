// Package _9_combination_sum - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* combination-sum - file brief introduce */
/*
modification history
-----------------------------
2020/4/19 10:00 PM, by gosyang, create
*/
/*
DESCRIPTION
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的数字可以无限制重复被选取。
说明：
所有数字（包括 target）都是正整数。
解集不能包含重复的组合。 
示例 1:
输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
示例 2:
输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
*/
package _9_combination_sum

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	// sort
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
		// left 的使用，每次从i开始，可以重复，但i前的不要了
		DFS(candidates, target-candidates[i], append(cur, candidates[i]), i, res)
	}
}


// 这里的关键是如何保证不是全排列，重复的可以去掉，那就是先排序，然后DFS的时候对于用过的值且确保多次重复后没用的就不循环了