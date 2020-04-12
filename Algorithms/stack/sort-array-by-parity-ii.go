// Package stack - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* sort-array-by-parity-ii - file brief introduce */
/*
modification history
-----------------------------
2020/4/12 6:14 PM, by gosyang, create
*/
/*
DESCRIPTION
给定一个非负整数数组 A， A 中一半整数是奇数，一半整数是偶数。

对数组进行排序，以便当 A[i] 为奇数时，i 也是奇数；当 A[i] 为偶数时， i 也是偶数。

你可以返回任何满足上述条件的数组作为答案。

输入：[4,2,5,7]
输出：[4,5,2,7]
解释：[4,7,2,5]，[2,5,4,7]，[2,7,4,5] 也会被接受。
*/
package stack

// 开辟了新的空间，且，可以在找出奇数就放在最终的结果上
func sortArrayByParityII(A []int) []int {
	l := len(A)
	odd := make([]int, l/2)
	m := 0
	even := make([]int, l/2)
	n := 0
	for i := 0; i < l; i++ {
		if A[i] % 2 == 0 {
			even[n] = A[i]
			n += 1
		} else {
			odd[m] = A[i]
			m += 1
		}
	}
	res := make([]int, l)
	for i := 0; i < l; i++ {
		if i % 2 == 0 {
			res[i] = even[i/2]
		} else {
			res[i] = odd[i/2]
		}
	}
	return res
}

// 不额外开辟空间，就直接找并替换就行
func sortArrayByParityII2(A []int) []int {
	j := 1
	for i := 0; i < len(A); i += 2 {
		// 偶数位的地方有奇数
		if A[i] % 2 != 0 {
			// 奇数位的地方找到偶数
			for A[j] % 2 == 1 {
				j += 2
			}
			A[i], A[j] = A[j], A[i]
		}
	}
	return A
}