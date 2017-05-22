#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: search_insert_position_test.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import search_insert_position


class SearchInsertPositionTestCase(unittest.TestCase):
    def setUp(self):
        self.solution = search_insert_position.Solution()

    def tearDown(self):
        self.solution = None

    def test_sample(self):
        self.assertEqual(self.solution.searchInsert([1, 3, 5, 6], 5), 2)
        self.assertEqual(self.solution.searchInsert([1, 3, 5, 6], 2), 1)
        self.assertEqual(self.solution.searchInsert([1, 3, 5, 6], 7), 4)
        self.assertEqual(self.solution.searchInsert([1, 3, 5, 6], 0), 0)
    
    def test_empty(self):
        self.assertEqual(self.solution.searchInsert([], 10), 0)
        
        
if __name__ == "__main__":
    unittest.main()


