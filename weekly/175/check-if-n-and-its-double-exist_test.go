// Package _75 - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* check-if-n-and-its-double-exist_test - file brief introduce */
/*
modification history
-----------------------------
2020/2/9 10:31 AM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _175

import "testing"

func TestCheckIfExist(t *testing.T) {
	tests := []struct {
		name string
		input []int
		expect bool
	}{
		{
			name: "2-no",
			input: []int{1,1},
			expect: false,
		},
		{
			name: "2-yes",
			input: []int{2,1},
			expect: true,
		},
		{
			name: "10,2,5,3",
			input: []int{10,2,5,3},
			expect: true,
		},
		{
			name: "7,1,14,11",
			input: []int{7,1,14,11},
			expect: true,
		},
		{
			name: "3,1,7,11",
			input: []int{3,1,7,11},
			expect: false,
		},
		{
			name: "-1,0,1",
			input: []int{-1,0,1},
			expect: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if get := checkIfExist(test.input); get != test.expect {
				t.Errorf("checkIfExist(%v)=%v != %v", test.input, get, test.expect)
			}
		})
	}
}


