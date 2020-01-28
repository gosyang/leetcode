// Package _0_climbing_stairs - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* climbing_stairs_test - file brief introduce */
/*
modification history
-----------------------------
2020/1/28 9:16 PM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _0_climbing_stairs

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		expect int
	}{
		{
			// 1
			name:   "1",
			input:  1,
			expect: 1,
		},
		{
			// 2
			name:   "2",
			input:  2,
			expect: 2,
		},
		{
			// 3
			name:   "3",
			input:  3,
			expect: 3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if get := climbStairs(test.input); get != test.expect {
				t.Errorf("climbStairs(%d)=%d != %d", test.input, get, test.expect)
			}
		})
	}
}
