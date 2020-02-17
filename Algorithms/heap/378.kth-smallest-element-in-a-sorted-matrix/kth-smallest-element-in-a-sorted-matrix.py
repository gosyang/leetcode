#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: kth-smallest-element-in-a-sorted-matrix
Author: gosyang
Date: 2020/2/17 10:10 AM
给定一个 n x n 矩阵，其中每行和每列元素均按升序排序，找到矩阵中第k小的元素。
请注意，它是排序后的第k小元素，而不是第k个元素。

你可以假设 k 的值永远是有效的, 1 ≤ k ≤ n2 。
"""
import heapq

class Solution(object):
    def kthSmallest(self, matrix, k):
        """
        :type matrix: List[List[int]]
        :type k: int
        :rtype: int
        """
        q = []
        n = len(matrix)
        for i in range(n):
            for j in range(n):
                m = matrix[i][j]
                if len(q) < k:
                    heapq.heappush(q, -m)
                else:
                    if -m < q[0]:
                        if j == 0:
                            return -heapq.heappop(q)
                        break
                    heapq.heapreplace(q, -m)
        return -heapq.heappop(q)

# 类似176 周赛第一题，二维矩阵有序的话，最左边的如果大了，那就不用再往后循环了