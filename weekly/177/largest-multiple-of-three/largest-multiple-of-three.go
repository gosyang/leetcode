// Package largest_multiple_of_three - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* largest-multiple-of-three - file brief introduce */
/*
modification history
-----------------------------
2020/2/23 11:45 AM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package largest_multiple_of_three

import "sort"

func largestMultipleOfThree(digits []int) string {
	sort.Sort(sort.Reverse(sort.IntSlice(digits)))

}
