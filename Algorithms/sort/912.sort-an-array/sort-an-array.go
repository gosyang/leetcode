// Package _12_sort_an_array - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* sort-an-array - file brief introduce */
/*
modification history
-----------------------------
2020/4/15 12:12 PM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _12_sort_an_array

import "math/rand"

func heapSort(nums []int) {
	len := len(nums) - 1
	buildMaxHeap(nums, len)
	for i := len; i >= 1; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		len -= 1
		maxHeapify(nums, 0, len)
	}
}

func buildMaxHeap(nums []int, len int) {
	for i := len / 2; i >= 0; i-- {
		maxHeapify(nums, i, len)
	}
}

func maxHeapify(nums []int, i, len int) {
	for (i<<1)+1 <= len {
		lson := i<<1 + 1
		rson := i<<1 + 2
		large := 0
		if lson <= len && nums[lson] > nums[i] {
			large = lson
		} else {
			large = i
		}
		if rson <= len && nums[rson] > nums[large] {
			large = rson
		}
		if i != large {
			nums[i], nums[large] = nums[large], nums[i]
			i = large
		} else {
			break
		}
	}
}

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, low, high int) {
	if low < high {
		randIndex := rand.Intn(high-low+1) + low
		nums[low], nums[randIndex] = nums[randIndex], nums[low]
		lowIndex := getIndex(nums, low, high)

		quickSort(nums, low, lowIndex-1)
		quickSort(nums, lowIndex+1, high)
	}
}

func getIndex(nums []int, low, high int) int {
	tmp := nums[low]
	for low < high {
		for low < high && nums[high] >= tmp {
			high -= 1
		}
		nums[low] = nums[high]
		for low < high && nums[low] <= tmp {
			low += 1
		}
		nums[high] = nums[low]
	}
	nums[low] = tmp
	return low
}

func bubleSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func sortArray2(nums []int) []int {
	mergeSort(nums, 0, len(nums)-1)
	return nums
}

func mergeSort(nums []int, l, r int) {
	m := l + (r-l)/2
	mergeSort(nums, l, m)
	mergeSort(nums, m+1, r)
	if nums[m] > nums[m+1] {
		merge(nums, l, m, r)
	}
}

func merge(nums []int, l, m, r int) {
	tmp := make([]int, r-l+1)
	i := l
	j := m + 1
	k := 0
	for ; k < r-l+1; k++ {
		if i > m {
			tmp[k] = nums[j]
			j++
		} else if j > r {
			tmp[k] = nums[i]
			i++
		} else if nums[i] < nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
	}
	for l := l; l <= r; l++ {
		nums[l] = tmp[l]
	}
}
