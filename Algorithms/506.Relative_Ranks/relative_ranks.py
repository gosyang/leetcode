#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: relative_ranks.py
Author: gosyang
Date: 2017/05/24 10:04:03
"""


class Solution(object):
    def findRelativeRanks(self, nums):
        """
        :type nums: List[int]
        :rtype: List[str]
        """
        # python 的方法由于封装的较好, 没有太多底层优化的空间
        # 除非手写 index 方法看用 桶原理 做
        before = nums[:]
        nums.sort(reverse=True)
        res = []
        for i in before:
            index = nums.index(i)
            if index == 0:
                res.append("Gold Medal")
            elif index == 1:
                res.append("Silver Medal")
            elif index == 2:
                res.append("Bronze Medal")
            else:
                res.append(str(index + 1))
        return res
                
                
                
        
        

