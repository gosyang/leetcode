// Package _ - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* 5340.count-negative-numbers-in-a-sorted-matrix - file brief introduce */
/*
modification history
-----------------------------
2020/5341.product-of-the-last-k-numbers/16 10:19 AM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _340_count_negative_numbers_in_a_sorted_matrix

func countNegatives(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	sum := 0
	for i := m-1; i >= 0; i-- {
		for j := n-1; j >=0; j-- {
			if grid[i][j] >= 0 {
				if j == n-1 {
					return sum
				}
				break
			}
			sum += 1
		}
	}
	return sum
}