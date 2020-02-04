// Package _38_range_sum_of_bst - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* range-sum-of-bst - file brief introduce */
/*
modification history
-----------------------------
2020/2/4 1:50 PM, by gosyang, create
*/
/*
DESCRIPTION
给定二叉搜索树的根结点 root，返回 L 和 R（含）之间的所有结点的值的和。

二叉搜索树保证具有唯一的值。
提示：

树中的结点数量最多为 10000 个。
最终的答案保证小于 2^31。
*/
package _38_range_sum_of_bst


// Definition for a binary tree node.
type TreeNode struct {
	Val int
    Left *TreeNode
    Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	if root.Val > R {
		return rangeSumBST(root.Left, L, R)
	}
	if root.Val < L {
		return rangeSumBST(root.Right, L, R)
	}
	return rangeSumBST(root.Left, L, R) + root.Val + rangeSumBST(root.Right, L, R)
}

