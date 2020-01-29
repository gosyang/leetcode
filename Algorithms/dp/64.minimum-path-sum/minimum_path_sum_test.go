// Package _4_minimum_path_sum - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* minimum_path_sum_test - file brief introduce */
/*
modification history
-----------------------------
2020/1/29 9:53 PM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _4_minimum_path_sum

import "testing"

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		name   string
		input  [][]int
		expect int
	}{
		{
			name:   "[[0]]",
			input:  [][]int{{0}},
			expect: 0,
		},
		{
			name:   "[[1]]",
			input:  [][]int{{1}},
			expect: 1,
		},
		{
			name:   "[[1,1],[1,1]]",
			input:  [][]int{{1, 1}, {1, 1}},
			expect: 3,
		},
		{
			name:   "[[1,3,1],[1,5,1],[4,2,1]]",
			input:  [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			expect: 7,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if get := minPathSum(test.input); get != test.expect {
				t.Errorf("minPathSum(%v)=%d != %d", test.input, get, test.expect)
			}
		})
	}
}
