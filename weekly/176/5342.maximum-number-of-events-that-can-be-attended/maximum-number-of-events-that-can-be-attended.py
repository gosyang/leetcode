#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: maximum-number-of-events-that-can-be-attended
Author: gosyang
Date: 2020/2/16 1:44 PM
给你一个数组 events，其中 events[i] = [startDayi, endDayi] ，表示会议 i 开始于 startDayi ，结束于 endDayi 。

你可以在满足 startDayi <= d <= endDayi 中的任意一天 d 参加会议 i 。注意，一天只能参加一个会议。

请你返回你可以参加的 最大 会议数目。
"""
import heapq

class Solution(object):
    def maxEvents(self, events):
        """
        :type events: List[List[int]]
        :rtype: int
        """
        ans = 0
        end = []
        # 开始时间由大到小
        events = sorted(events, reverse=True)
        for i in range(1, 100001):
            while events and events[-1][0] == i:
                heapq.heappush(end, events.pop()[1])
            while end and end[0] < i:
                heapq.heappop(end)
            if end:
                heapq.heappop(end)
                ans += 1
        return ans
