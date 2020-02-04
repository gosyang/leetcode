// Package _00_search_in_a_binary_search_tree - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* search-in-a-binary-search-tree - file brief introduce */
/*
modification history
-----------------------------
2020/2/4 3:51 PM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _00_search_in_a_binary_search_tree

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if val == root.Val {
		return root
	}
	if val < root.Val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}

// 一次AC