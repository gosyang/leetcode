#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: student_attendance_record_I.py
Author: gosyang
Date: 2017/07/01 16:06:38
"""

class Solution(object):
    def checkRecord(self, s):
        """
        :type s: str
        :rtype: bool
        """
        if s is None:
            return True
        # 连续的两个LL，那就是必须得有LLL
        if s.count("A") > 1 or "LLL" in s:
            return False
        return True
