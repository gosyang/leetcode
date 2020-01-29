// Package _98_house_robber - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* house_robber_test - file brief introduce */
/*
modification history
-----------------------------
2020/1/29 12:50 PM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _98_house_robber

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			name:   "1",
			input:  []int{1},
			expect: 1,
		},
		{
			name:   "0",
			input:  []int{},
			expect: 0,
		},
		{
			name:   "[1,2,3,1]",
			input:  []int{1, 2, 3, 1},
			expect: 4,
		},
		{
			name:   "[2,1,1,2]",
			input:  []int{2, 1, 1, 2},
			expect: 4,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if get := rob(test.input); get != test.expect {
				t.Errorf("rob(%v)=%d != %d", test.input, get, test.expect)
			}
		})
	}
}
