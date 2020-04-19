// Package face38_zi_fu_chuan_de_pai_lie_lcof - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* zi-fu-chuan-de-pai-lie-lcof - file brief introduce */
/*
modification history
-----------------------------
2020/4/19 12:32 PM, by gosyang, create
*/
/*
DESCRIPTION
输入一个字符串，打印出该字符串中字符的所有排列。
你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。

示例:

输入：s = "abc"
输出：["abc","acb","bac","bca","cab","cba"]
*/
package face38_zi_fu_chuan_de_pai_lie_lcof

func permutation(s string) []string {
	res := []string{}
	dfs([]byte(s), 0, &res)
	return res
}
// 这里处理字符数组的技巧！！
func dfs(s []byte, start int, res *[]string) {
	if start == len(s)-1 {
		*res = append(*res, string(s))
		return
	}
	// 这里的bug点在于 给的字符可能是有重复的，排列出来是可能重复的，保证个字符在某个位置一次
	m := map[byte]bool{}
	for i := start; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			continue
		}
		m[s[i]] = true
		s[i], s[start] = s[start], s[i]
		dfs(s, start+1, res)
		// 换回来
		s[i], s[start] = s[start], s[i]
	}
}