// Package _02_binary_tree_level_order_traversal - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* binary-tree-level-order-traversal - file brief introduce */
/*
modification history
-----------------------------
2020/4/18 4:39 PM, by gosyang, create
*/
/*
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
*/
package _02_binary_tree_level_order_traversal

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		tmp := []int{}

		for i := 0; i < size; i++ {
			tmp = append(tmp, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[size:]
		res = append(res, tmp)
	}
	return res
}

// 最简单的如果结果是一维数组，就直接每次tmp的地方res append就行。关键是这里是批量bfs用熟