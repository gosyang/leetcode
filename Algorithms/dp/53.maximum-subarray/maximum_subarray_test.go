// Package _3_maximum_subarray - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* maximum_subarray_test - file brief introduce */
/*
modification history
-----------------------------
2020/1/28 9:34 PM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _3_maximum_subarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			name:   "[1]",
			input:  []int{1},
			expect: 1,
		},
		{
			name:   "[2,3,-4]",
			input:  []int{2, 3, -4},
			expect: 5,
		},
		{
			name:   "[-2,1,-3,4,-1,2,1,-5,4]",
			input:  []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expect: 6,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if get := maxSubArray2(test.input); get != test.expect {
				t.Errorf("maxSubArray(%v)=%d != %d", test.input, get, test.expect)
			}
		})
	}
}
