#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: intersection_of_two_arrays_II.py
Author: baidu(baidu@baidu.com)
Date: 2017/05/20 23:20:03
"""
import collections


class Solution(object):
    def intersect(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: List[int]
        """
        # 没有一次AC的原因, 理解题意错误, 以为是最长公共子序列,
        # 其实是按出现次数统计然后列出来就行了, 不需要有序
        # 假如内存不够, 那就一点一点读好了, 统计维护好就行
        a, b = map(collections.Counter, (nums1, nums2))
        return list((a & b).elements())

