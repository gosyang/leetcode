#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: the-masseuse-lcci
Author: gosyang
Date: 2020/5/31 5:33 PM
"""
class Solution(object):
    def massage(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        一个有名的按摩师会收到源源不断的预约请求，每个预约都可以选择接或不接。在每次预约服务之间要有休息时间，因此她不能接受相邻的预约。
        给定一个预约请求序列，替按摩师找到最优的预约集合（总预约时间最长），返回总的分钟数。

        输入： [2,1,4,5,3,1,1,3]
        输出： 12
        解释： 选择 1 号预约、 3 号预约、 5 号预约和 8 号预约，总时长 = 2 + 4 + 3 + 3 = 12。

        思路一 时间空间都不太好
        dp[i]代表选i的情况下的最长时间
        dp[i] = max{dp[j]}+nums[i], 0<=j<i-1
        最终求max{dp[i]}, 0<=i<len(nums)
        dp[0] = nums[0]

        """
        l = len(nums)
        if l == 0:
            return 0
        dp = [0] * l
        dp[0] = nums[0]
        res = dp[0]
        for i in range(1, l):
            m = 0
            for j in range(i-1):
                m = max(m, dp[j])
            dp[i] = m + nums[i]
            res = max(res, dp[i])
        return res

    def massage2(self, nums):
        """
        思路二
        第i个不接
        dp[i][0]=max(dp[i−1][0],dp[i−1][1])
        第i个接
        dp[i][1]=dp[i−1][0]+nums
        最终看
        max(dp[n][0],dp[n][1])
        """
        n = len(nums)
        if n == 0:
            return 0

        dp0, dp1 = 0, nums[0]
        for i in range(1, n):
            tdp0 = max(dp0, dp1)   # 计算 dp[i][0]
            tdp1 = dp0 + nums[i]   # 计算 dp[i][1]
            dp0, dp1 = tdp0, tdp1

        return max(dp0, dp1)
