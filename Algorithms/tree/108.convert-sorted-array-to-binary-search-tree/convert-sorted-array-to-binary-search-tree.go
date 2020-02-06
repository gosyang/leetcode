// Package convert_sorted_array_to_binary_search_tree - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* convert-sorted-array-to-binary-search-tree - file brief introduce */
/*
modification history
-----------------------------
2020/2/5 2:17 PM, by gosyang, create
*/
/*
DESCRIPTION
将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
*/
package convert_sorted_array_to_binary_search_tree

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	mid := n / 2
	return &TreeNode{
		Left: sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
		Val: nums[mid],
	}
}

// 这道题属于简单的，升序数组组成平衡的bst，那就是去mid的元素做根即可
// 类似的题目，一起研究

