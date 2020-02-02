// Package _0_regular_expression_matching - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* regular_expression_matching - file brief introduce */
/*
modification history
-----------------------------
2020/2/2 9:20 AM, by gosyang, create
*/
/*
DESCRIPTION
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
*/
package _0_regular_expression_matching

func isMatch(s string, p string) bool {
	sl := len(s)
	pl := len(p)
	dp := make([][]bool, sl+1)
	for i := 0; i <= sl; i++ {
		dp[i] = make([]bool, pl+1)
	}

	dp[0][0] = true
	// 如果p是空一切白搭就是false默认了，直接判断s是空
	for j := 1; j <= pl; j++ {
		//偶数位是*, 可以消掉
		if j&1 == 0 && p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i <= sl; i++ {
		for j := 1; j <= pl; j++ {
			if (s[i-1] != p[j-1]) && p[j-1] != '*' && p[j-1] != '.' {
				dp[i][j] = false
			} else if p[j-1] != '*' && p[j-1] != '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if j >= 2 {
				if p[j-2] == s[i-1] || p[j-2] == '.' {
					dp[i][j] = dp[i-1][j] || dp[i][j-2]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}
	return dp[sl][pl]
}

// 两个字符串遍历的问题，存在最优子结构，dp
// a. dp[i][j] 代表s的i位，p的j位结尾的两个能否匹配(i,j从1开始)
// b. dp[i][j] = 1. dp[i-1][j-1], s[i-1] == p[j-1] 且p[j-1]不是特殊字符
//			     2. false, s[i-1] != p[j-1] 且p[j-1]不是特殊字符
// 				 3. !!重要 dp[i-1][j]|| dp[i][j-2], p[j-1] == * 且 (p[j-1] == s[i-1] || p[j-2] == .)
// 			     4. dp[i][j-2], p[j-1] == * 且 (p[j-2] != s[i-1] && p[j-2]!=.)(这里要考虑*是0个的情况)
// 			     5. dp[i-1][j-1], p[j-1] == .
// c. dp[0][0] == true(两个空串), dp[0][j is 偶数] = dp[0][j-2]
// d. dp[len(s)][len(p)]
// 因为* 可以0个或多个，如果有*，且匹配上了（情况3）a ab*a*举例，最后一个*判断的时候，不一定使用到了这个a*，可能前面就达成了
// dp[i-1][j] 表示 a* 可能匹配了s的1个或多个, 看s的前一个
// dp[i][j-2] 表示 a* 没用, 就可以匹配
