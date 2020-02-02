// Package _39_word_break - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* word_break_test - file brief introduce */
/*
modification history
-----------------------------
2020/2/2 8:47 PM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _39_word_break

import (
	"testing"
)

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name   string
		inputS string
		inputD []string
		expect bool
	}{
		{
			name:   "leetcode",
			inputS: "leetcode",
			inputD: []string{"leet", "code"},
			expect: true,
		},
		{
			name:   "applepenapple",
			inputS: "applepenapple",
			inputD: []string{"apple", "pen"},
			expect: true,
		},
		{
			name:   "catsandog",
			inputS: "catsandog",
			inputD: []string{"cats", "dog", "sand", "and"},
			expect: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if get := wordBreak(test.inputS, test.inputD); get != test.expect {
				t.Errorf("wordBreak(%s, %v)=%v != %v", test.inputS, test.inputD, get, test.expect)
			}
		})
	}
}
