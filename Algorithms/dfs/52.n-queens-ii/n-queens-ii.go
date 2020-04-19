// Package _2_n_queens_ii - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* n-queens-ii - file brief introduce */
/*
modification history
-----------------------------
2020/4/19 9:37 PM, by gosyang, create
*/
/*
DESCRIPTION
n 皇后，输出解法个数
*/
package _2_n_queens_ii

func totalNQueens(n int) int {
	if n < 1 {
		return 0
	}
	shu := map[int]bool{}
	pie := map[int]bool{}
	na := map[int]bool{}

	res := 0
	DFS(n, 0, shu, pie, na, []int{}, &res)
	return res
}

func DFS(n int, row int, shu, pie, na map[int]bool, cur []int, res *int) {
	if row >= n {
		// 这里需要特别注意要copy一份！！
		*res += 1
		return
	}
	for col := 0; col < n; col++ {
		if shu[col] || pie[row+col] || na[row-col] {
			continue
		}
		shu[col] = true
		na[row-col] = true
		pie[row+col] = true
		DFS(n, row+1, shu, pie, na, append(cur, col), res)
		shu[col] = false
		na[row-col] = false
		pie[row+col] = false
	}
}
