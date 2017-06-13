#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: reverse_vowels_of_a_string.py
Author: gosyang 
Date: 2017/06/03 13:31:54
"""


class Solution(object):
    def reverseVowels(self, s):
        """
        前后指针法
        :type s: str
        :rtype: str
        """
        vowel = 'AEIOUaeiou'
        # 还可以这么操作!
        s = list(s)
        i, j = 0, len(s) - 1
        while i < j:
            while s[i] not in vowel and i < j:
                i += 1
            while s[j] not in vowel and i < j:
                j -= 1
            s[i], s[j] = s[j], s[i]
            i, j = i + 1, j - 1
        return ''.join(s)
    
    def reverseVowels_timeout(self, s):
        """
        超时算法! 一个case通不过, O（2n）
        :type s: str
        :rtype: str
        """
        if len(s) == 0:
            return s
        volwels = set("aeiouAEIOU")
        volwels_index = []
        # enumerate 好方便
        for (index, value) in enumerate(s):
            if value in volwels:
                volwels_index.append(index)
        if len(volwels_index) == 0:
            return s
        reverse_index = volwels_index[::-1]
        # 如果是没有元音, 这里取第0个就错了!
        res = s[:volwels_index[0]]
        for (index, value) in enumerate(reverse_index):
            res += s[value]
            if index != len(reverse_index) - 1:
                res += s[volwels_index[index] + 1: volwels_index[index + 1]]
            else:
                # 非元音结尾
                res += s[volwels_index[-1] + 1:]
        return res
