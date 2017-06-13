#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: island_perimeter.py
Author: gosyang 
Date: 2017/06/04 15:43:56
"""


class Solution(object):
    def islandPerimeter(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        res = 0
        row = len(grid)
        if row == 0:
            return res
        column = len(grid[0])
        for i in range(row):
            for j in range(column):
                if grid[i][j] == 0:
                    continue
                if j == 0:
                    res += 1
                elif grid[i][j - 1] == 0:
                    res += 1
                if j == column - 1:
                    res += 1
                elif grid[i][j + 1] == 0:
                    res += 1
                
                if i == 0:
                    res += 1
                elif grid[i - 1][j] == 0:
                    res += 1
                if i == row - 1:
                    res += 1
                elif grid[i + 1][j] == 0:
                    res += 1
        return res
                    
                    
                
                    
                    
