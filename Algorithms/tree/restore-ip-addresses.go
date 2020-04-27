// Package tree - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* restore-ip-addresses - file brief introduce */
/*
modification history
-----------------------------
2020/4/21 8:19 PM, by gosyang, create
*/
/*
DESCRIPTION
给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

示例:

输入: "25525511135"
输出: ["255.255.11.135", "255.255.111.35"]
*/
package tree

import "strconv"

func restoreIpAddresses(s string) []string {
	l := len(s)
	res := []string{}
	if l < 4 || l > 12 {
		return res
	}
	dfs(s, 0, "", &res)
	return res
}

// dfs
func dfs(s string, n int, cur string, res *[]string) {
	if n == 3 {
		if len(s) == 0 || len(s) > 3 {
			return
		}
		sn, _ := strconv.Atoi(s)
		// 考虑 "0x" 和 "0xx"/"00x" 都是非法的
		if len(s) == 1 || len(s) == 2 && sn >= 10 || len(s) == 3  && sn >= 100 && sn <= 255 {
			*res = append(*res, cur+"."+s)
		}
		return
	}
	l := len(s)
	for i := 1; i <= min(l, 3); i++ {
		if i == 2 {
			sn, _ := strconv.Atoi(s[:2])
			if sn < 10 {
				continue
			}
		}
		if i == 3 {
			sn, _ := strconv.Atoi(s[:3])
			if  sn > 255 || sn < 100 {
				continue
			}
		}
		if n == 0 {
			dfs(s[i:], 1, s[:i], res)
		} else {
			dfs(s[i:], n+1, cur+"."+s[:i], res)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
