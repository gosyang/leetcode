// Package main - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* getDecimalValue_test - file brief introduce */
/*
modification history
-----------------------------
2019/12/19 2:35 PM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package main

import "testing"

func TestGetDecimalValue(t *testing.T) {
	tests := []struct {
		name   string
		input  *ListNode
		expect int
	}{
		{
			name: "0",
			input: &ListNode{
				Val: 0,
			},
			expect: 0,
		},
		{
			name: "1",
			input: &ListNode{
				Val: 1,
			},
			expect: 1,
		},
		{
			name: "01",
			input: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 1,
				},
			},
			expect: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if get := getDecimalValue(test.input); get != test.expect {
				t.Errorf("getDecimalValue want %d but get %d", test.expect, get)
			}
		})
	}
}
