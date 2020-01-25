// Package _36_lowest_common_ancestor_of_a_binary_tree - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* lowest-common-ancestor-of-a-binary-tree - file brief introduce */
/*
modification history
-----------------------------
2020/1/25 9:38 AM, by gosyang, create
*/
/*
DESCRIPTION
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

*/
package _36_lowest_common_ancestor_of_a_binary_tree

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	} else {
		return left
	}
}
// 递归的思路，