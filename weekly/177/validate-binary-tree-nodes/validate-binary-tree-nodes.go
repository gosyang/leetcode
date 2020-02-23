// Package validate_binary_tree_nodes - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* validate-binary-tree-nodes - file brief introduce */
/*
modification history
-----------------------------
2020/2/23 10:46 AM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package validate_binary_tree_nodes

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	m := map[int]bool{0: true}
	maxl := 0
	maxr := 0
	for i := 0; i < n; i++ {
		if leftChild[i] != -1 {
			if leftChild[i] <= maxl {
				return false
			}
			if _, ok := m[leftChild[i]]; !ok {
				m[leftChild[i]] = true
			} else {
				return false
			}
		}
		if rightChild[i] != -1 {
			if rightChild[i] <= maxr {
				return false
			}
			if _, ok := m[rightChild[i]]; !ok {
				m[rightChild[i]] = true
			} else {
				return false
			}
		}
	}
	for i := 1; i < n; i++ {
		if _, ok := m[i]; !ok {
			return false
		}
	}
	return true
}

// 比赛的时候假设编号从0开始一层一层递增的，因此如果出现非递增的肯定不行
// 如果出现重复的也不行，最后全部扫一遍发现有没有访问到的（除了0）也不行
