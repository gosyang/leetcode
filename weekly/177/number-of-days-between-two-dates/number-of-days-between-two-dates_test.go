// Package number_of_days_between_two_dates - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* number-of-days-between-two-dates_test - file brief introduce */
/*
modification history
-----------------------------
2020/2/23 10:39 AM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package number_of_days_between_two_dates

import "testing"

func TestDaysBetweenDates(t *testing.T) {
	tests := []struct {
		name   string
		input1 string
		input2 string
		expect int
	}{
		{
			name:   "0",
			input1: "2019-06-29",
			input2: "2019-06-29",
			expect: 0,
		},
		{
			name:   "1",
			input1: "2019-06-29",
			input2: "2019-06-30",
			expect: 1,
		},
		{
			name:   "2",
			input1: "2020-01-15",
			input2: "2019-12-31",
			expect: 15,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if get := daysBetweenDates(test.input1, test.input2); get != test.expect {
				t.Errorf("(%s, %s)=%d != %d", test.input1, test.input2, get, test.expect)
			}
		})
	}
}
