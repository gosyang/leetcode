// Package _76_argest_perimeter_triangle - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* argest-perimeter-triangle - file brief introduce */
/*
modification history
-----------------------------
2020/4/12 6:29 PM, by gosyang, create
*/
/*
DESCRIPTION
给定由一些正数（代表长度）组成的数组 A，返回由其中三个长度组成的、面积不为零的三角形的最大周长。

如果不能形成任何面积不为零的三角形，返回 0。
*/
package _76_argest_perimeter_triangle

import "sort"

func largestPerimeter(A []int) int {
	sort.Ints(A)
	for i := len(A)-3; i >= 0; i-- {
		if A[i] + A[i+1] > A[i+2] {
			return A[i] + A[i+1] + A[i+2]
		}
	}
	return 0
}

// 排序，从最大的来，如果第二第三大的相加也比他小，那再小的其实会更小，也就不用再枚举了