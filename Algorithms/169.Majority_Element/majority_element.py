#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: Hamming_Distance.py
Author: gosyang
Date: 2017/06/23 16:06:38
"""
import collections


class Solution(object):
    def majorityElement(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
	# 其他方法是排序后取len(nums)/2的元素
        return collections.Counter(nums).most_common(1)[0][0]
