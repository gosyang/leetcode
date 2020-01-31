// Package _1_decode_ways - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* decode_ways - file brief introduce */
/*
modification history
-----------------------------
2020/1/30 12:12 PM, by gosyang, create
*/
/*
DESCRIPTION
一条包含字母 A-Z 的消息通过以下方式进行了编码：

'A' -> 1
'B' -> 2
...
'Z' -> 26
给定一个只包含数字的非空字符串，请计算解码方法的总数。
*/
package _1_decode_ways

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	two := 1
	one := 1
	for i := 1; i < len(s); i++ {
		tmp := one
		// 如果最后一位是0，那就不能算i-1的所有解法了，只能算i-2的
		if s[i] == '0' {
			one = 0
		}
		if s[i-1] == '1' || s[i-1] == '2' && s[i] <= '6'{
			one += two
		}
		two = tmp
	}
	return one
}

// a. dp[i] 到第i个字符的总数
// b. dp[i] = dp[i-1] + dp[i-2]
// 			  1. 如果第i位是0， 则dp[i] = dp[i-2]
// 			  2. i-1和i 不在10-26内(01-09都是非法的), dp[i] = dp[i-1]
// c.边界条件 dp[0] = 1
