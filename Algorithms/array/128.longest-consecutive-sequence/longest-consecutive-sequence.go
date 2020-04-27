/*
给定一个未排序的整数数组，找出最长连续序列的长度。

要求算法的时间复杂度为 O(n)。

示例:

输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
*/
package main

// 找到起始的数字，然后递增, 找出长度
func longestConsecutive(nums []int) int {
	m := map[int]bool{}
	for _, n := range nums {
		m[n] = true
	}
	// 注意为空的情况
	longest := 0
	for k := range m {
		if _, ok := m[k-1]; ok {
			continue
		}
		l := 1
		for {
			if !m[k+1] {
				break
			}
			l++
			k++
		}
		if l > longest {
			longest = l
		}
	}
	return longest
}
