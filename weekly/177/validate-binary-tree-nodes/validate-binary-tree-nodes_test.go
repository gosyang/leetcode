// Package validate_binary_tree_nodes - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* validate-binary-tree-nodes_test - file brief introduce */
/*
modification history
-----------------------------
2020/2/23 11:19 AM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package validate_binary_tree_nodes

import "testing"

func TestValidateBinaryTreeNodes(t *testing.T) {
	tests := []struct {
		name   string
		inputN int
		inputL []int
		inputR []int
		expect bool
	}{
		{
			name:   "1",
			inputN: 1,
			inputL: []int{-1},
			inputR: []int{-1},
			expect: true,
		},
		{
			name:   "4",
			inputN: 4,
			inputL: []int{1, -1, 3, -1},
			inputR: []int{2, -1, -1, -1},
			expect: true,
		},
		{
			name:   "circle",
			inputN: 4,
			inputL: []int{1, -1, 3, -1},
			inputR: []int{2, 3, -1, -1},
			expect: false,
		},
		{
			name:   "r2",
			inputN: 2,
			inputL: []int{1, 0},
			inputR: []int{-1, -1},
			expect: false,
		},
		{
			name:   "6",
			inputN: 6,
			inputL: []int{1, -1, -1, 4, -1, -1},
			inputR: []int{2, -1, -1, 5, -1, -1},
			expect: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if get := validateBinaryTreeNodes(test.inputN, test.inputL, test.inputR); get != test.expect {
				t.Errorf("(%d, %v, %v)=%v != %v", test.inputN, test.inputL, test.inputR, get, test.expect)
			}
		})
	}
}
