#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: longest_uncommon_subsequence_II.py
Author: baidu(baidu@baidu.com)
Date: 2017/06/06 11:35:57
"""
import re


class Solution(object):
    def checkinSetSub(self, s, word):
        pattern = re.compile(".*" + '.*'.join([i for i in word]) + ".*")
        for w in s:
            match = pattern.search(w)
            if match:
                return True
        return False
            
    def findLUSlength(self, strs):
        """
        :type strs: List[str]
        :rtype: int
        """
        strs.sort(lambda x, y: len(y) - len(x))
        res = -1
        longest = set()
        before = set()
        for word in strs:
            if word not in longest:
                # 避免把重复的再算一遍
                if word not in before:
                    longest.add(word)
                if res == -1:
                    if self.checkinSetSub(before, word):
                        # print "check [{}] insub [{}]".format(word, before)
                        before.add(word)
                        continue
                    before.add(word)
                    res = len(word)
                else:
                    # 假如不同但是又长度相同, 需要加入到set里
                    before.add(word)
                    # print "add {} to {}".format(word, before)
                    if len(word) < res:
                        return res
            else:
                longest.remove(word)
                # print "delete {}, now {}".format(word, longest)
                if len(longest) == 0:
                    res = -1
        return res
