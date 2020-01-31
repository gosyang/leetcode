// Package _2_edit_distance - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* edit_distance_test - file brief introduce */
/*
modification history
-----------------------------
2020/1/31 3:13 PM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _2_edit_distance

import "testing"

func TestMinDistance(t *testing.T) {
	tests := []struct{
		name string
		input1 string
		input2 string
		expect int
	}{
		{
			name: "empty",
			input1: "",
			input2: "",
			expect: 0,
		},
		{
			name: "same",
			input1: "a",
			input2: "a",
			expect: 0,
		},
		{
			name: "one replace",
			input1: "a",
			input2: "b",
			expect: 1,
		},
		{
			name: "delete",
			input1: "ab",
			input2: "b",
			expect: 1,
		},
		{
			name: "add",
			input1: "b",
			input2: "ab",
			expect: 1,
		},
		{
			name: "horse-ros",
			input1: "horse",
			input2: "ros",
			expect: 3,
		},
		{
			name: "intention-execution",
			input1: "intention",
			input2: "execution",
			expect: 5,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if get := minDistance(test.input1, test.input2); get != test.expect {
				t.Errorf("minDistance(%s, %s)=%d != %d", test.input1, test.input2, get, test.expect)
			}
		})
	}
}