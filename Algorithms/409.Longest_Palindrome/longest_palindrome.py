#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: longest_palindrome.py
Author: gosyang
Date: 2017/06/29 16:06:38
"""

import collections

class Solution(object):
    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: int
        """
        m = collections.Counter(s).most_common()

        res = 0
        jishu_has_add = False
        for i in m:
            if i[1] % 2 != 0:
                if not jishu_has_add:
                    res += i[1]
                    jishu_has_add = True
                else:
                    res += i[1] - 1
            else:
                res += i[1]
        return res
