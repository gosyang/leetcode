// Package _0_valid_parentheses - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* valid-parentheses - file brief introduce */
/*
modification history
-----------------------------
2020/1/19 11:25 AM, by gosyang, create
*/
/*
DESCRIPTION
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。
*/
package _0_valid_parentheses

func isValid(s string) bool {
	stack := make([]rune, 0)
	parentMap := map[rune]rune{'(':')', '[':']', '{':'}'}

	for _, v := range s {
		if _, ok := parentMap[v]; ok {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			// 如果栈顶的元素不可以匹配
			if parentMap[stack[len(stack)-1]] != v {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}

// 这个题使用rune做单字符是go的收获！使用golang的字符串遍历