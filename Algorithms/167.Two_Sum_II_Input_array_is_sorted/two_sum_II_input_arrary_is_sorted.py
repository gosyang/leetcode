#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: two_sum_II_input_arrary_is_sorted.py
Author: gosyang
Date: 2017/06/14 13:57:39
"""


class Solution(object):
    def twoSum(self, numbers, target):
        """
        Given an array of integers that is already sorted in ascending order,
        find two numbers such that they add up to a specific target number.

        The function twoSum should return indices of the two numbers
        such that they add up to the target, where index1 must be less than index2.
        Please note that your returned answers (both index1 and index2) are not zero-based.

        You may assume that each input would have exactly one solution
        and you may not use the same element twice.

        Input: numbers={2, 7, 11, 15}, target=9
        Output: index1=1, index2=2
        
        :type numbers: List[int]
        :type target: int
        :rtype: List[int]
        """
        # O(n^2) 的方法, 加了强剪枝可以过
        lens = len(numbers)
        for index1 in range(lens - 1):
            # 不加这个会超时, 因为存在唯一解, 假如上一个index1 算过没有,那么下一个相同的就不用算了
            if index1 >= 1 and numbers[index1] == numbers[index1 - 1]:
                continue
            diff = target - numbers[index1]
            for index2 in range(index1 + 1, lens):
                if numbers[index2] == diff:
                    return [index1 + 1, index2 + 1]
                if numbers[index2] > diff:
                    break

    def twoSumOn(self, numbers, target):
        # 或者两个指针的方法
        index1 = 0
        index2 = len(numbers) - 1
        while index1 < index2:
            sum = numbers[index1] + numbers[index2]
            if sum == target:
                return [index1 + 1, index2 + 1]
            elif sum > target:
                index2 -= 1
            else:
                index1 += 1
