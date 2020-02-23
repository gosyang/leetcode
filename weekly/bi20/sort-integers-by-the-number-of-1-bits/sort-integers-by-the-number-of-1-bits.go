// Package sort_integers_by_the_number_of_1_bits - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* sort-integers-by-the-number-of-1-bits - file brief introduce */
/*
modification history
-----------------------------
2020/2/22 10:32 PM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package sort_integers_by_the_number_of_1_bits

import "sort"

type MapSorter []Item

type Item struct {
	Key int
	Val int
}

func (ms MapSorter) Len() int {
	return len(ms)
}

func (ms MapSorter) Less(i, j int) bool {
	if ms[i].Val == ms[j].Val {
		return ms[i].Key < ms[j].Key
	}
	return ms[i].Val < ms[j].Val
}

func (ms MapSorter) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func sortByBits(arr []int) []int {
	m := MapSorter{}
	for i := 0; i < len(arr); i++ {
		m = append(m, Item{
			Key: arr[i],
			Val: oneNums(arr[i]),
		})
	}
	sort.Sort(m)
	res := make([]int, 0)
	for i := 0; i < len(m); i++ {
		res = append(res, m[i].Key)
	}
	return res
}

func oneNums(num int) int {
	count := 0
	for count = 0; num != 0; num >>= 1 {
		if (num & 1) != 0 {
			count += 1
		}
	}
	return count
}
