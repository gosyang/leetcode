#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: time_out.py
Author: baidu(baidu@baidu.com)
Date: 2017/05/21 22:19:01
"""
import collections


class Solution(object):
    def distance(self, p1, p2):
        """

        :param p1: list[int]
        :param p2: list[int]
        :return:   double
        """
        x = p1[0] - p2[0]
        y = p1[1] - p2[1]
        return x * x + y * y
    
    def numberOfBoomerangs(self, points):
        """
        :type points: List[List[int]]
        :rtype: int
        """
        res = 0
        dis_map = []
        for i in range(len(points)):
            l = []
            for j in range(i):
                l.append(dis_map[j][i])
            l.append(-1)
            for j in range(i + 1, len(points)):
                l.append(self.distance(points[i], points[j]))
            dis_map.append(l)
        # print dis_map
        
        # 超时在于此处计算浪费时间, Counter估计耗时长
        for i in dis_map:
            counter = collections.Counter(i)
            for k in counter:
                if counter[k] > 1:
                    res += counter[k] * (counter[k] - 1)
        return res


