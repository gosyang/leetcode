// Package __longest_palindromic_substring - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* longest_palindromicsubstring_test - file brief introduce */
/*
modification history
-----------------------------
2020/1/31 12:44 PM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package __longest_palindromic_substring

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct{
		name string
		input string
		expect string
	}{
		{
			name: "empty string",
			input: "",
			expect: "",
		},
		{
			name: "a",
			input: "a",
			expect: "a",
		},
		{
			name: "ac",
			input: "ac",
			expect: "a",
		},
		{
			name: "cbbd",
			input: "cbbd",
			expect: "bb",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			if get := longestPalindrome(test.input); get != test.expect {
				t.Errorf("longestPalindrome(%s)=%s != %s", test.input, get, test.expect)
			}
		})
	}
}