// Package sliding_window_maximum - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* sliding-window-maximum - file brief introduce */
/*
modification history
-----------------------------
2020/4/18 9:59 PM, by gosyang, create
*/
/*
DESCRIPTION
给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回滑动窗口中的最大值。
*/
package sliding_window_maximum

func maxSlidingWindow(nums []int, k int) []int {
	max := []int{}
	res := []int{}
	if len(nums) == 0 {
		return res
	}

	slide := nums[:k]
	for i := 0; i < k; i++ {
		if len(max) == 0 || nums[i] > max[0] {
			max = []int{nums[i]}
		} else {
			for j := 1; j < len(max); j++ {
				if nums[i] > max[j] {
					max[j] = nums[i]
					max = max[:j+1]
				}
			}
			max = append(max, nums[i])
		}
	}
	res = append(res, max[0])

	for i := k; i < len(nums); i++ {
		del := slide[0]
		slide = append(slide[1:], nums[i])
		if del == max[0] {
			max = max[1:]
		}
		for j := 0; j < len(max); j++ {
			if nums[i] > max[j] {
				max = max[:j]
				break
			}
		}
		max = append(max, nums[i])
		res = append(res, max[0])
	}

	return res
}