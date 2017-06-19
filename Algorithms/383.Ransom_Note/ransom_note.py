#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: ransom_note.py
Author: gosyang 
Date: 2017/06/19 09:42:59
"""
import collections


class Solution(object):
    def canConstruct(self, ransomNote, magazine):
        """
        :type ransomNote: str
        :type magazine: str
        :rtype: bool
        """
        note = collections.Counter(ransomNote)
        mag = collections.Counter(magazine)
        
        for i in note:
            if note[i] > mag[i]:
                return False
        return True
    
        # 字符串有个好方法, str.count("A") 显示有几个, 因此可以先对note 取set,然后同样判断


