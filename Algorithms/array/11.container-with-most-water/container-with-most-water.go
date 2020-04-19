// Package _1_container_with_most_water - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* container-with-most-water - file brief introduce */
/*
modification history
-----------------------------
2020/4/19 6:15 PM, by gosyang, create
*/
/*
DESCRIPTION
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。
*/
package _1_container_with_most_water

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	max := -1
	for i < j {
		now := min(height[i], height[j]) * (j - i)
		if now > max {
			max = now
		}
		// 双指针，移动每次比较小的那个
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}