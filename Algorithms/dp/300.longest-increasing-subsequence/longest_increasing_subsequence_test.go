// Package _00_longest_increasing_subsequence - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* longest_increasing_subsequence_test - file brief introduce */
/*
modification history
-----------------------------
2020/2/3 1:09 PM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _00_longest_increasing_subsequence

import "testing"

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			name:   "empty",
			input:  []int{},
			expect: 0,
		},
		{
			name:   "1",
			input:  []int{1},
			expect: 1,
		},
		{
			name:   "1,2",
			input:  []int{1, 2},
			expect: 2,
		},
		{
			name:   "3,2,1",
			input:  []int{3, 2, 1},
			expect: 1,
		},
		{
			name:   "10,9,2,5,3,7,101,18",
			input:  []int{10, 9, 2, 5, 3, 7, 101, 18},
			expect: 4,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if get := lengthOfLIS(test.input); get != test.expect {
				t.Errorf("lengthOfLIS(%v)=%d != %d", test.input, get, test.expect)
			}
		})
	}
}
