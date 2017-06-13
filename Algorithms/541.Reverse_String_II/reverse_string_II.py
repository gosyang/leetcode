#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: reverse_string_II.py
Author: gosyang 
Date: 2017/06/02 09:39:31
"""


class Solution(object):
    def reverseStr(self, s, k):
        """
        :type s: str
        :type k: int
        :rtype: str
        """
        res = ""
        for i in range(len(s) / (2 * k)):
            res += s[i * 2 * k: i * 2 * k + k][::-1] + s[i * 2 * k + k: (i + 1) * 2 * k]
        left_length = len(s) % (2 * k)
        # 易错点1: 假如正好是2k, 则后面不能再算, 否则切片容易出错 s[-0:] 这种 ==> 单测未覆盖
        if left_length == 0:
            return res
        # 易错点2: 假如是k的长度, 需要按照小于k算,否则 -left_length + k 超出 ==> 单测未覆盖
        # 因此最好还是 在单测的时候把这种可能的情况穷举!! 区间和分界点!
        if left_length <= k:
            res += s[-left_length:][::-1]
        else:
            res += s[-left_length: -left_length + k][::-1] + s[-left_length + k:]
        return res

