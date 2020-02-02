// Package _0_regular_expression_matching - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* regular_expression_matching_test - file brief introduce */
/*
modification history
-----------------------------
2020/2/2 9:20 AM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _0_regular_expression_matching

import "testing"

func TestIsMatch(t *testing.T) {
	tests := []struct {
		name   string
		inputS string
		inputP string
		expect bool
	}{
		{
			name:   "both empty",
			inputS: "",
			inputP: "",
			expect: true,
		},
		{
			name:   "S empty",
			inputS: "",
			inputP: "a",
			expect: false,
		},
		{
			name:   "S empty, P .*",
			inputS: "",
			inputP: ".*",
			expect: true,
		},
		{
			name:   "S not empty, P .*",
			inputS: "abc",
			inputP: ".*",
			expect: true,
		},
		{
			name:   "aa,a",
			inputS: "aa",
			inputP: "a",
			expect: false,
		},
		{
			name:   "aa,a*",
			inputS: "aa",
			inputP: "a*",
			expect: true,
		},
		{
			name:   "c,cd*",
			inputS: "c",
			inputP: "cd*",
			expect: true,
		},
		{
			name:   "aab,c*a*b",
			inputS: "aab",
			inputP: "c*a*b",
			expect: true,
		},
		{
			name:   "mississippi",
			inputS: "mississippi",
			inputP: "mis*is*p*.",
			expect: false,
		},
		{
			name:   "aaa,ab*a*c*a",
			inputS: "aaa",
			inputP: "ab*a*c*a",
			expect: true,
		},
		{
			name:   "a,.*..a*",
			inputS: "a",
			inputP: ".*..a*",
			expect: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if get := isMatch(test.inputS, test.inputP); get != test.expect {
				t.Errorf("isMatch(%s, %s)=%v != %v", test.inputS, test.inputP, get, test.expect)
			}
		})
	}
}
