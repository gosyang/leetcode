// Package _04_Maximum_Depth_of_Binary_Tree - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* maximum-depth-of-binary-tree - file brief introduce */
/*
modification history
-----------------------------
2020/4/18 4:52 PM, by gosyang, create
*/
/*
DESCRIPTION
树的最大深度，用树的bfs做
*/
package _04_Maximum_Depth_of_Binary_Tree

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[size:]
		res++
	}
	return res
}