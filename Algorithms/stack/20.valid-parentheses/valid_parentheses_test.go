// Package _0_valid_parentheses - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* valid_parentheses_test - file brief introduce */
/*
modification history
-----------------------------
2020/1/19 11:26 AM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _0_valid_parentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		input string
		expect bool
	}{
		{
			name: "()",
			input: "()",
			expect: true,
		},
		{
			name: "()[]",
			input: "()[]",
			expect: true,
		},
		{
			name: "(]",
			input: "(]",
			expect: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if get := isValid(test.input); get != test.expect {
				t.Errorf("isValid(%s)= %v != %v, fail", test.input, get, test.expect)
			}
		})
	}
}
