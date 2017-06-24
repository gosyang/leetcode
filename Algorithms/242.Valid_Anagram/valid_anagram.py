#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: valid_anagram.py
Author: gosyang
Date: 2017/06/24 16:06:38
"""

import collections


class Solution(object):
    def isAnagram(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        # 提升速度的话就是先判断长度，再累计，在过程中发现不一样的字母就return
        return collections.Counter(s) == collections.Counter(t)
        
