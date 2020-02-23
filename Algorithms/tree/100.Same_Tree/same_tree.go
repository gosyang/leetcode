// Package _00_Same_Tree - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* same_tree - file brief introduce */
/*
modification history
-----------------------------
2020/2/4 11:30 AM, by gosyang, create
*/
/*
DESCRIPTION
给定两个二叉树，编写一个函数来检验它们是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
*/
package _00_Same_Tree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}

// 递归遍历，O（N)
