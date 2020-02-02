// Package main - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* unique_binary_search_trees_test - file brief introduce */
/*
modification history
-----------------------------
2020/1/31 5:17 PM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package main

import "testing"

func TestNumTrees(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		expect int
	}{
		{
			name:   "1",
			input:  1,
			expect: 1,
		},
		{
			name:   "2",
			input:  2,
			expect: 2,
		},
		{
			name:   "3",
			input:  3,
			expect: 5,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if get := numTrees(test.input); get != test.expect {
				t.Errorf("numTrees(%d)=%d != %d", test.input, get, test.expect)
			}
		})
	}
}
