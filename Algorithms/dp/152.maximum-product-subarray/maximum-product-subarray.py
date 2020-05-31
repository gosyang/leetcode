#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: maximum-product-subarray
Author: gosyang
Date: 2020/5/30 9:33 PM

给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

[2,3,-2,4]
[1]
[-2,0,-1]
[-2,3,-4]
[7,-2,-4]
"""
class Solution(object):
    def maxProduct(self, nums):
        """
        :type nums: List[int]
        :rtype: int

        关键就是 要保存正的最大值和负的最小值
        """
        res = nums[0]
        dp = res
        mdp = res if res < 0 else 0
        for i in range(1, len(nums)):
            if nums[i] >= 0:
                dp = max(nums[i], dp*nums[i])
                mdp *= nums[i]
            else:
                dp, mdp = max(nums[i]*mdp, nums[i]), min(nums[i]*dp, nums[i])
            res = max(res, dp)
        return res
