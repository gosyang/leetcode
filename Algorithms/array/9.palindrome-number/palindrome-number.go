// Package __palindrome_number - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* palindrome-number - file brief introduce */
/*
modification history
-----------------------------
2020/5/10 11:51 PM, by gosyang, create
*/
/*
DESCRIPTION
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

负数不是
*/
package __palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	y := x
	n := []int{}
	for y > 0 {
		n = append(n, y%10)
		y = y/10
	}
	l := len(n)
	for i := 0; i < l/2; i++ {
		if n[i] != n[l-i-1] {
			return false
		}
	}
	return true
}
// 时间O(lgn) 空间O(lgn)

// 优化，翻转一半，我们将原始数字除以 10，然后给反转后的数字乘上 10，所以，当原始数字小于反转后的数字时，就意味着我们已经处理了一半位数的数字
// 可以优化到空间O(1)

