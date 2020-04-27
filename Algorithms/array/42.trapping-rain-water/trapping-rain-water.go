// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
package main

func trap(height []int) int {
	l := 0
	r := len(height) - 1
	ans := 0
	if r < 0 {
		return 0
	}
	lmax := height[l]
	rmax := height[r]
	for l < r {
		if lmax > rmax {
			r--
			if height[r] < rmax {
				ans += rmax - height[r]
			} else {
				rmax = height[r]
			}
		} else {
			l++
			if height[l] < lmax {
				ans += lmax - height[l]
			} else {
				lmax = height[l]
			}
		}
	}
	return ans
}

// 思想就是 左右指针，左边找到最大的，那右边的小于右最大的就可以盛水了
