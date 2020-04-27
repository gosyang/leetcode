/*
给定一个未经排序的整数数组，找到最长且连续的的递增序列。

示例 1:

输入: [1,3,5,4,7]
输出: 3
解释: 最长连续递增序列是 [1,3,5], 长度为3。
尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为5和7在原数组里被4隔开。
示例 2:

输入: [2,2,2,2,2]
输出: 1
解释: 最长连续递增序列是 [2], 长度为1。
*/
package main

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ml := 1
	// 这里max 得是1，否则如果没有递增的话，结果是0不是期望的1
	max := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			ml++
			// 这里不能加在下面，可能是纯单调的
			if ml > max {
				max = ml
			}
		} else {
			ml = 1
		}
	}
	return max
}
