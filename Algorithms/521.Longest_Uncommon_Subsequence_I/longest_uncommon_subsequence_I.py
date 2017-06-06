#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: longest_uncommon_subsequence_I.py
Author: baidu(baidu@baidu.com)
Date: 2017/06/06 10:58:20
"""


class Solution(object):
    def findLUSlength(self, a, b):
        """
        最长非公共子序列, 如果不相等, 必然是整体作为子序列, 两个更长的那个
        :type a: str
        :type b: str
        :rtype: int
        """
        return a != b and max(len(a), len(b)) or -1
