#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: number_of_boomeranges.py
Author: gosyang
Date: 2017/05/21 21:16:32
"""


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

    def put_into_map(self, map, x):
        if x not in map:
            map[x] = 0
        map[x] += 1
        return map
    
    def numberOfBoomerangs(self, points):
        """
        :type points: List[List[int]]
        :rtype: int
        """
        res = 0
        dis_map = []
        for i in range(len(points)):
            l = []
            s = {}
            for j in range(i):
                l.append(dis_map[j][i])
                # 可以简化为 s[dis_map[j][i]] = s.get(dis_map[j][i], 0) + 1
                self.put_into_map(s, dis_map[j][i])
            l.append(-1)
            for j in range(i + 1, len(points)):
                distance = self.distance(points[i], points[j])
                l.append(distance)
                self.put_into_map(s, distance)
            dis_map.append(l)
            # 如果是1, 减1 就是0 乘下也没什么, TODO 到底哪个费时间?
            for k in s:
                if s[k] > 1:
                    res += s[k] * (s[k] - 1)
        # print dis_map
        
        return res
    
    # 同样的思路的简洁版, 将距离保存, 使用get(default = 0) 来简化判断key存在不存在!
    # res = 0
    # for p in points:
    #     cmap = {}
    #     for q in points:
    #         f = p[0] - q[0]
    #         s = p[1] - q[1]
    #         cmap[f * f + s * s] = 1 + cmap.get(f * f + s * s, 0) 赞!
    #     for k in cmap:
    #         res += cmap[k] * (cmap[k] - 1)
    # return res
                    
            
            
        
            
        

