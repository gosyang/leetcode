// Package closest_divisors - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* closest-divisors_test - file brief introduce */
/*
modification history
-----------------------------
2020/2/23 11:40 AM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package closest_divisors

import "testing"

func TestClosestDivisors(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		expect []int
	}{
		{
			name:   "1",
			input:  1,
			expect: []int{1, 2},
		},
		{
			name:   "8",
			input:  8,
			expect: []int{3, 3},
		},
		{
			name:   "999",
			input:  999,
			expect: []int{40, 25},
		},
		{
			name:   "123",
			input:  123,
			expect: []int{5, 25},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if get := closestDivisors(test.input); len(get) == len(test.expect) {
				t.Errorf("(%d)=%v != %v", test.input, get, test.expect)
			}
		})
	}
}
