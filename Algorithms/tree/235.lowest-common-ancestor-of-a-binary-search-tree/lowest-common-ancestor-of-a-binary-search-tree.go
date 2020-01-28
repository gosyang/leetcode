// Package _35_lowest_common_ancestor_of_a_binary_search_tree - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* lowest-common-ancestor-of-a-binary-search-tree - file brief introduce */
/*
modification history
-----------------------------
2020/1/25 9:48 AM, by gosyang, create
*/
/*
DESCRIPTION
二叉搜索树，返回p和q的最近公共祖先
*/
package _35_lowest_common_ancestor_of_a_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

// root判断不能少，养成习惯，有指针有成员变量，先判断指针nil
// 如果是q，p的大小不知道p一定比q小，那就放在最后了

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}

// 非递归更快, 总结tree的递归非递归的关系，非递归就是root的更新！
