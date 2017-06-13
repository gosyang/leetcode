#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: detect_capital.py
Author: gosyang 
Date: 2017/06/08 10:27:59
"""


class Solution(object):
    def detectCapitalUse(self, word):
        """
        :type word: str
        :rtype: bool
        """
        if len(word) == 1:
            return True
        first_capital = word[0].isupper()
        others_all_capital = word[1:].isupper()
        others_all_lower = word[1:].islower()
    
        return first_capital and others_all_lower or \
               first_capital and others_all_capital or \
               not first_capital and others_all_lower

        # Python istitle() 方法检测字符串中所有的单词拼写首字母是否为大写，且其他字母为小写。
        # return  word.isupper() or word.istitle() or word.islower()