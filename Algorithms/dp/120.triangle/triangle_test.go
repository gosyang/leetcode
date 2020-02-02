// Package _20_triangle - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* triangle_test - file brief introduce */
/*
modification history
-----------------------------
2020/2/2 7:54 PM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _20_triangle

import (
	"testing"
)

func TestMinimumTotal(t *testing.T) {
	tests := []struct {
		name   string
		input  [][]int
		expect int
	}{
		{
			name:   "empty",
			input:  [][]int{},
			expect: 0,
		},
		{
			name:   "1",
			input:  [][]int{{1}},
			expect: 1,
		},
		{
			name:   "example",
			input:  [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
			expect: 11,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if get := minimumTotal(test.input); get != test.expect {
				t.Errorf("minimumTotal(%v)=%d != %d", test.input, get, test.expect)
			}
		})
	}
}
