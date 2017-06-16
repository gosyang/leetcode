#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: assign_cookies.py
Author: gosyang 
Date: 2017/06/16 09:30:30
"""


class Solution(object):
    def findContentChildren(self, g, s):
        """
        Assume you are an awesome parent and want to give your children some cookies.
        But, you should give each child at most one cookie.
        Each child i has a greed factor gi, which is the minimum size of a cookie
        that the child will be content with; and each cookie j has a size sj.
        If sj >= gi, we can assign the cookie j to the child i, and the child i will be content.
        
        Your goal is to maximize the number of your content children and output the maximum number.
        
        :type g: List[int]
        :type s: List[int]
        :rtype: int
        """
        # 先排序, 用最小的饼干满足最小的孩子
        # tips: 下面的特殊判断可以从 159ms -> 92ms, 没道理吧。。
        if not g or not s:
            return 0
        res, i, j = 0, 0, 0
        g.sort()
        s.sort()
        while i < len(g) and j < len(s):
            if g[i] <= s[j]:
                i += 1
                res += 1
            j += 1
        return res
