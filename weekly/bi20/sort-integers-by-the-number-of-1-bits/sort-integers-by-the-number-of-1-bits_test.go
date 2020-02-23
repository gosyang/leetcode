// Package sort_integers_by_the_number_of_1_bits - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* sort-integers-by-the-number-of-1-bits_test - file brief introduce */
/*
modification history
-----------------------------
2020/2/22 10:36 PM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package sort_integers_by_the_number_of_1_bits

import "testing"

func TestSortByBits(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect []int
	}{
		{
			name:   "1",
			input:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			expect: []int{0, 1, 2, 4, 8, 3, 5, 6, 7},
		},
		{
			name:   "2",
			input:  []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1},
			expect: []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024},
		},
		{
			name:   "3",
			input:  []int{10000, 10000},
			expect: []int{10000, 10000},
		},
		{
			name:   "4",
			input:  []int{2, 3, 5, 7, 11, 13, 17, 19},
			expect: []int{2, 3, 5, 17, 7, 11, 13, 19},
		},
		{
			name:   "5",
			input:  []int{10, 100, 1000, 10000},
			expect: []int{10, 100, 10000, 1000},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if get := sortByBits(test.input); len(get) == len(test.expect) {
				t.Errorf("sortByBits(%v)=%v != %v", test.input, get, test.expect)
			}
		})
	}
}
