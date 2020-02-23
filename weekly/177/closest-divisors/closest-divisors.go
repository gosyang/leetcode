// Package closest_divisors - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* closest-divisors - file brief introduce */
/*
modification history
-----------------------------
2020/2/23 11:25 AM, by gosyang, create
*/
/*
DESCRIPTION
给你一个整数 num，请你找出同时满足下面全部要求的两个整数：

两数乘积等于  num + 1 或 num + 2
以绝对差进行度量，两数大小最接近
你可以按任意顺序返回这两个整数。
*/
package closest_divisors

import "math"

func closestDivisors(num int) []int {
	num1 := num + 1
	num2 := num + 2
	sqrt1 := int(math.Sqrt(float64(num1)))
	sqrt2 := int(math.Sqrt(float64(num2)))
	res := make([]int, 0)
	diff := math.MaxInt32
	for i := sqrt1; i >= 1; i-- {
		if num1%i == 0 {
			other := num1 / i
			if other-i < diff {
				res = []int{i, other}
				diff = other - i
			} else if other-i == diff {
				res = append(res, i)
				res = append(res, other)
			}
			break
		}
	}
	for i := sqrt2; i >= 1; i-- {
		if num2%i == 0 {
			other := num2 / i
			if other-i < diff {
				res = []int{i, other}
				diff = other - i
			} else if other-i == diff {
				res = append(res, i)
				res = append(res, other)
			}
			break
		}
	}
	return res
}
