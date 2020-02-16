#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: last-stone-weight
Author: gosyang
Date: 2020/2/16 5:08 PM
有一堆石头，每块石头的重量都是正整数。

每一回合，从中选出两块最重的石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：

如果 x == y，那么两块石头都会被完全粉碎；
如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
最后，最多只会剩下一块石头。返回此石头的重量。如果没有石头剩下，就返回 0。
提示：

1 <= stones.length <= 30
1 <= stones[i] <= 1000
"""
import heapq
class Solution(object):
    def lastStoneWeight(self, stones):
        """
        :type stones: List[int]
        :rtype: int
        """
        if len(stones) == 1:
            return stones[0]

        stones = [-i for i in stones]
        heapq.heapify(stones)
        while len(stones) > 1:
            y = heapq.heappop(stones)
            x = heapq.heappop(stones)
            if x == y:
                continue
            else:
                heapq.heappush(stones, y-x)
        if len(stones) == 1:
            return 0 - stones[0]
        else:
            return 0

# python的是小顶堆，但凡用到大顶堆的就反着放，并最终结果记得反回来



