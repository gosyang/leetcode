// Package _1_n_queens - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* n-queens - file brief introduce */
/*
modification history
-----------------------------
2020/4/19 9:32 PM, by gosyang, create
*/
/*
DESCRIPTION
n皇后
*/
package _1_n_queens

func solveNQueens(n int) [][]string {
	if n < 1 {
		return [][]string{[]string{}}
	}
	shu := map[int]bool{}
	pie := map[int]bool{}
	na := map[int]bool{}

	res := [][]int{}
	DFS(n, 0, shu, pie, na, []int{}, &res)
	return output(n, res)
}

func DFS(n int, row int, shu, pie, na map[int]bool, cur []int, res *[][]int) {
	if row >= n {
		// 这里需要特别注意要copy一份！！
		tmp := make([]int, n)
		copy(tmp, cur)
		*res = append(*res, tmp)
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

func output(n int, res [][]int) [][]string {
	ans := [][]string{}
	for i := 0; i < len(res); i++ {
		a := []string{}
		r := res[i]
		for j := 0; j < len(r); j++ {
			s := ""
			for k := 0; k < r[j]; k++ {
				s += "."
			}
			s += "Q"
			for k := 0; k < n-r[j]-1; k++ {
				s += "."
			}
			a = append(a, s)
		}
		ans = append(ans, a)
	}
	return ans
}