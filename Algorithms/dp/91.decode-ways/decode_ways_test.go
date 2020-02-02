// Package _1_decode_ways - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* decode_ways_test - file brief introduce */
/*
modification history
-----------------------------
2020/1/30 12:12 PM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _1_decode_ways

import "testing"

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "0",
			input:  "0",
			expect: 0,
		},
		{
			name:   "00",
			input:  "00",
			expect: 0,
		},
		{
			name:   "000",
			input:  "000",
			expect: 0,
		},
		{
			name:   "101",
			input:  "101",
			expect: 1,
		},
		{
			name:   "1",
			input:  "1",
			expect: 1,
		},
		{
			name:   "10",
			input:  "10",
			expect: 1,
		},
		{
			name:   "12",
			input:  "12",
			expect: 2,
		},
		{
			name:   "226",
			input:  "226",
			expect: 3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if get := numDecodings(test.input); get != test.expect {
				t.Errorf("numDecodings(%s)=%d != %d", test.input, get, test.expect)
			}
		})
	}
}
