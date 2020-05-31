#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: range-sum-query-immutable
Author: gosyang
Date: 2020/5/31 9:31 PM
给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。

示例：

给定 nums = [-2, 0, 3, -5, 2, -1]，求和函数为 sumRange()

sumRange(0, 2) -> 1
sumRange(2, 5) -> -1
sumRange(0, 5) -> -3
说明:

你可以假设数组不可变。
会多次调用 sumRange 方法
"""

class NumArray(object):

    def __init__(self, nums):
        """
        :type nums: List[int]
        说白了就是先计算好0-j, 减去0-i就行，
        状态方程设计个dp l+1个可以简单一点
        """
        self.dp = [0] * (len(nums)+1)
        for i in range(len(nums)):
            self.dp[i+1] = self.dp[i] + nums[i]

    def sumRange(self, i, j):
        """
        :type i: int
        :type j: int
        :rtype: int
        """
        return self.dp[j+1]-self.dp[i]



# Your NumArray object will be instantiated and called as such:
# obj = NumArray(nums)
# param_1 = obj.sumRange(i,j)