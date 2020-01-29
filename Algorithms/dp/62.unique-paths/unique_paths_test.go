// Package _2_unique_paths - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* unique_paths_test - file brief introduce */
/*
modification history
-----------------------------
2020/1/29 10:42 PM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _2_unique_paths

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name   string
		inputM int
		inputN int
		expect int
	}{
		{
			name:   "1,1",
			inputM: 1,
			inputN: 1,
			expect: 1,
		},
		{
			name:   "3,2",
			inputM: 3,
			inputN: 2,
			expect: 3,
		},
		{
			name:   "7,3",
			inputM: 7,
			inputN: 3,
			expect: 28,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if get := uniquePaths(test.inputM, test.inputN); get != test.expect {
				t.Errorf("uniquePaths(%d, %d)=%d != %d", test.inputM, test.inputN, get, test.expect)
			}
		})
	}
}
