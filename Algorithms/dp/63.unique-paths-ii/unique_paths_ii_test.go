// Package _3_unique_paths_ii - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* unique_paths_ii_test - file brief introduce */
/*
modification history
-----------------------------
2020/1/30 10:58 AM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _3_unique_paths_ii

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	tests := []struct {
		name   string
		input  [][]int
		expect int
	}{
		{
			name:   "[[0]]",
			input:  [][]int{{0}},
			expect: 1,
		},
		{
			name:   "[[1]]",
			input:  [][]int{{1}},
			expect: 0,
		},
		{
			name:   "[[1,0]]",
			input:  [][]int{{1, 0}},
			expect: 0,
		},
		{
			name:   "[[1,1],[1,1]]",
			input:  [][]int{{1, 1}, {1, 1}},
			expect: 0,
		},
		{
			name:   "[[0,0,0],[0,1,0],[0,0,0]]",
			input:  [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			expect: 2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if get := uniquePathsWithObstacles(test.input); get != test.expect {
				t.Errorf("uniquePathsWithObstacles(%v)=%d != %d", test.input, get, test.expect)
			}
		})
	}
}
