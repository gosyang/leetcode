// Package one - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* one_test - file brief introduce */
/*
modification history
-----------------------------
2020/2/16 10:21 AM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package one

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		name string
		input [][]int
		expect int
	}{
		{
			name: "one",
			input: [][]int{{-1}},
			expect: 1,
		},
		{
			name: "two",
			input: [][]int{{1,-1},{-1,-1}},
			expect: 3,
		},
		{
			name: "three",
			input: [][]int{{3,2},{1,0}},
			expect: 0,
		},
		{
			name: "four",
			input: [][]int{{4,3,2,-1},{3,2,1,-1},{1,1,-1,-2},{-1,-1,-2,-3}},
			expect: 8,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if get := countNegatives(test.input); get != test.expect {
				t.Errorf("countNe(%v)=%d != %d", test.input, get, test.expect)
			}
		})
	}
}
