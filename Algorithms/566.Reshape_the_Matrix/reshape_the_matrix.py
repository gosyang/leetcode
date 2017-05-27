#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: reshape_the_matrix.py
Author: gosyang
Date: 2017/05/27 13:33:06
"""


class Solution(object):
    def matrixReshape(self, nums, r, c):
        """
        :type nums: List[List[int]]
        :type r: int
        :type c: int
        :rtype: List[List[int]]
        """
        if len(nums) * len(nums[0]) != r * c:
            return nums
        res = []
        res_line = []
        for line in nums:
            for colum in line:
                res_line.append(colum)
                if len(res_line) == c:
                    res.append(res_line)
                    res_line = []
        return res

