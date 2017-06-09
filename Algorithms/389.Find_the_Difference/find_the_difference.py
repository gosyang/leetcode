#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: find_the_difference.py
Author: gosyang
Date: 2017/06/09 21:00:54
"""

class Solution(object):
    def findTheDifference(self, s, t):
        """
        先排序, 再顺序比较获得, 丑 55ms
        :type s: str
        :type t: str
        :rtype: str
        """
        s, t = list(s), list(t)
        s.sort()
        t.sort()
        for i in range(min(len(s), len(t))):
            if s[i] != t[i]:
                if len(s) < len(t):
                    return t[i]
                else:
                    return s[i]
        if len(s) < len(t):
            return t[-1]
        else:
            return s[-1]
