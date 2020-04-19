// Package _13_path_sum_ii - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* path-sum-ii - file brief introduce */
/*
modification history
-----------------------------
2020/4/18 4:10 PM, by gosyang, create
*/
/*
DESCRIPTION
一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。

说明: 叶子节点是指没有子节点的节点。
*/
package _13_path_sum_ii

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
func pathSum(root *TreeNode, sum int) [][]int {
	// 用递归
	if root == nil {
		return nil
	}
	var res [][]int
	dfs(root, sum, []int{}, &res)
	return res
}

// 辅助函数，arr是当前dfs前的arr是多少
func dfs(root *TreeNode, sum int, arr []int, res *[][]int) {
	if root == nil {
		return
	}
	n := sum - root.Val
	arr = append(arr, root.Val)
	if n == 0 && root.Left == nil && root.Right == nil {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		*res = append(*res, tmp)
	}
	dfs(root.Left, n, arr, res)
	dfs(root.Right, n, arr, res)

	arr = arr[:len(arr)-1]
}