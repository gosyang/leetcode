// Package _75 - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* minimum-number-of-steps-to-make-two-strings-anagram_test - file brief introduce */
/*
modification history
-----------------------------
2020/2/9 10:45 AM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _175

import (
	"testing"
)

func TestMinSteps(t *testing.T) {
	tests := []struct {
		name   string
		inputS string
		inputT string
		expect int
	}{
		{
			name:   "empty",
			inputS: "",
			inputT: "",
			expect: 0,
		},
		{
			name:   "1",
			inputS: "bab",
			inputT: "aba",
			expect: 1,
		},
		{
			name:   "2",
			inputS: "leetcode",
			inputT: "practice",
			expect: 5,
		},
		{
			name:   "anagram",
			inputS: "anagram",
			inputT: "mangaar",
			expect: 0,
		},
		{
			name:   "xxyyzz",
			inputS: "xxyyzz",
			inputT: "xxyyzz",
			expect: 0,
		},
		{
			name:   "friend",
			inputS: "friend",
			inputT: "family",
			expect: 4,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if get := minSteps(test.inputS, test.inputT); get != test.expect {
				t.Errorf("minSteps(%s, %s)=%d != %d", test.inputS, test.inputT, get, test.expect)
			}
		})
	}
}
