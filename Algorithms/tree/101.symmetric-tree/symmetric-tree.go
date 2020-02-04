// Package _01_symmetric_tree - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* symmetric-tree - file brief introduce */
/*
modification history
-----------------------------
2020/2/4 2:10 PM, by gosyang, create
*/
/*
DESCRIPTION
给定一个二叉树，检查它是否是镜像对称的。
*/
package _01_symmetric_tree

// Definition for a binary tree node.
type TreeNode struct {
	Val int
    Left *TreeNode
    Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)
}

// 1. 中序遍历是对称的即可 O(n)
// 2. 递归子树是左右对称的即可，辅助函数