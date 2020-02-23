// Package _65_univalued_binary_tree - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* univalued-binary-tree - file brief introduce */
/*
modification history
-----------------------------
2020/2/4 2:47 PM, by gosyang, create
*/
/*
DESCRIPTION
如果二叉树每个节点都具有相同的值，那么该二叉树就是单值二叉树。

只有给定的树是单值二叉树时，才返回 true；否则返回 false。
*/
package _65_univalued_binary_tree

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Right == nil {
		return root.Val == root.Left.Val && isUnivalTree(root.Left)
	}
	if root.Left == nil {
		return root.Val == root.Right.Val && isUnivalTree(root.Right)
	}
	return root.Val == root.Left.Val && isUnivalTree(root.Left) && root.Val == root.Right.Val && isUnivalTree(root.Right)
}

func isUnivalTreeGood(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := root.Left == nil || root.Val == root.Left.Val && isUnivalTree(root.Left)
	right := root.Right == nil || root.Val == root.Right.Val && isUnivalTree(root.Right)
	return left && right
}
